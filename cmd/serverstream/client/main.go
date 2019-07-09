package main

import (
	pb "github.com/noguchi-hiroshi/grpc-test/proto/serverstream"
	"github.com/noguchi-hiroshi/grpc-test/serverstream"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(serverstream.URL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial %v\n", err)
	}
	defer conn.Close()
	sc := pb.NewServerStreamClient(conn)
	client := serverstream.NewClient(sc)
	greet(client)
}

func greet(c serverstream.Client) {
	if err := c.Greet(1, 10); err != nil {
		log.Fatalln(err)
	}
}
