package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/gorilla/mux"
	"github.com/nawajar/duck/configuration"
	"github.com/nawajar/duck/facebook"
)

const (
	defaultPort   = "8000"
	defaultAppEnv = "LOCAL"
)

func main() {
	var (
		appPort  = getEnvString("PORT", defaultPort)
		httpAddr = flag.String("http.addr", ":"+appPort, "HTTP listen address")
	)

	flag.Parse()

	log.Println("The service is starting...")
	config := configuration.New()
	router := mux.NewRouter()
	log.Println("The service is starting... " + config.AppURL)
	log.Println("===")
	
	facebookHandler := facebook.MakeHandler()

	r := router.NewRoute().Subrouter()
	{
		 r.HandleFunc("/status", facebookHandler.Hello).Methods("GET")
	}

	server := httpServer(*httpAddr, router)

	startServer(server)

	waitingForSignal(os.Interrupt, syscall.SIGTERM)

	log.Println("The service is shutting down...")

	forceShutdownAfter(server, time.Second*30)

	log.Println("terminated...")

	os.Exit(0)
}

func forceShutdownAfter(server *http.Server, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	server.Shutdown(ctx)
}

func waitingForSignal(sig ...os.Signal) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, sig...)

	s := <-stop
	log.Println("Got signal ", s.String())
}

func startServer(server *http.Server) {
	ch := make(chan error, 1)

	go func() {
		ch <- server.ListenAndServe()
	}()

	select {
	case err := <-ch:
		log.Fatal(err)
	default:
		log.Println("The service is ready to listen and serve.")
	}
}

func httpServer(httpAddr string, router http.Handler) *http.Server {
	return &http.Server{
		Addr: httpAddr,
		// WriteTimeout: time.Second * 15,
		// ReadTimeout:  time.Second * 15,
		// IdleTimeout:  time.Second * 60,
		Handler: router,
	}
}

func getEnvString(env, fallback string) string {
	result := os.Getenv(env)
	if result == "" {
		return fallback
	}
	return result
}


