package main

import (
	"fmt"
	"github.com/noguchi-hiroshi/grpc-test/rest"
	"time"
)

func main() {
	r := rest.NewClient()
	start := time.Now()
	for i := 0; i < 1000; i++ {
		r.Get()
	}
	fmt.Println(time.Since(start))
}
