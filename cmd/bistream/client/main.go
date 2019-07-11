package main

import (
	"github.com/noguchi-hiroshi/grpc-test/bistream"
	pb "github.com/noguchi-hiroshi/grpc-test/proto/bistream"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(bistream.URL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial %v\n", err)
	}
	defer conn.Close()
	sc := pb.NewBiStreamClient(conn)
	client := bistream.NewClient(sc)
	greet(client)
}

func greet(c bistream.Client) {
	if err := c.Greet(); err != nil {
		log.Fatalln(err)
	}
}
