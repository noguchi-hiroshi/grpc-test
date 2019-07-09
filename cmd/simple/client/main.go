package main

import (
	"fmt"
	pb "github.com/noguchi-hiroshi/grpc-test/proto/simple"
	"github.com/noguchi-hiroshi/grpc-test/simple"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial(simple.URL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial %v\n", err)
	}
	defer conn.Close()
	sc := pb.NewSimpleClient(conn)
	client := simple.NewClient(sc)
	bench(client)
}

func bench(c simple.Client) {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		c.Bench()
	}
	fmt.Println(time.Since(start))
}

func greet(c simple.Client) {
	if err := c.Greet(); err != nil {
		log.Fatalln(err)
	}
}
