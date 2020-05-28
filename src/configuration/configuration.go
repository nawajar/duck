package configuration

import (
	"os"
	"reflect"
)

type Configuration struct {
	PORT   string `env:"PORT"`
	AppURL string `env:"APP_URL" default:"http://localhost:8000"`
}

func New() Configuration {
	conf := Configuration{}
	v := reflect.ValueOf(&conf).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		envKey := fieldType.Tag.Get("env")
		envValue, ok := os.LookupEnv(envKey)

		switch ok {
		case true:
			field.SetString(envValue)
		case false:
			field.SetString(fieldType.Tag.Get("default"))
		}
	}

	return conf
}
