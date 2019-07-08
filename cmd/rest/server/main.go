package main

import (
	"github.com/noguchi-hiroshi/grpc-test/src/rest"
	"net/http"
)

func main() {
	r := rest.NewHandler()
	http.HandleFunc("/", r.Handler)
	http.ListenAndServe(rest.PORT, nil)
}
