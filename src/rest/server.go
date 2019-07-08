package rest

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

type handler struct {
}

func NewHandler() Handler {
	return &handler{}
}

func (h *handler) Handler(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Response{Message: "Hello"})
	if err != nil {
		log.Fatal(err)
	}
	w.Write(b)
}
