package facebook

import (
	"fmt"
	"net/http"
)

//MakeHandler maybe we have to make it first
func MakeHandler() *Handler {
	return &Handler{}
}

//Handler sure this's handler
type Handler struct {
}

//Hello just hello...
func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "OK!")
}