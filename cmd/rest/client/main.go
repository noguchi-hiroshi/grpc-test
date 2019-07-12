package main

import (
	"fmt"
	"github.com/noguchi-hiroshi/grpc-test/rest"
	"time"
)

func main() {
	r := rest.NewClient()
	bench(r, 10)
	bench(r, 100)
	bench(r, 300)
	bench(r, 500)
	bench(r, 1000)
	bench(r, 3000)
	bench(r, 5000)
	bench(r, 10000)
}

func bench(r rest.Client, max int) {
	start := time.Now()
	for i := 0; i < max; i++ {
		r.Get()
	}
	fmt.Println(time.Since(start))
}
