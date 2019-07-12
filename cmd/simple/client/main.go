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
	bench(client, 10)
	bench(client, 100)
	bench(client, 300)
	bench(client, 500)
	bench(client, 1000)
	bench(client, 3000)
	bench(client, 5000)
	bench(client, 10000)
}

func bench(c simple.Client, max int) {
	start := time.Now()
	for i := 0; i < max; i++ {
		c.Bench()
	}
	fmt.Println(time.Since(start))
}

func greet(c simple.Client) {
	if err := c.Greet(); err != nil {
		log.Fatalln(err)
	}
}
