package main

import (
	"github.com/noguchi-hiroshi/grpc-test/src/rest"
)

func main() {
	r := rest.NewClient()
	r.Get()
}
