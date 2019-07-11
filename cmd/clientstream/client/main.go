package main

import (
	pb "github.com/noguchi-hiroshi/grpc-test/proto/clientstream"
	"github.com/noguchi-hiroshi/grpc-test/clientstream"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(clientstream.URL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial %v\n", err)
	}
	defer conn.Close()
	sc := pb.NewClientStreamClient(conn)
	client := clientstream.NewClient(sc)
	greet(client)
}

func greet(c clientstream.Client) {
	if err := c.Greet(); err != nil {
		log.Fatalln(err)
	}
}
