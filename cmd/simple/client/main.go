package main

import (
	pb "github.com/noguchi-hiroshi/grpc-test/proto/simple"
	"github.com/noguchi-hiroshi/grpc-test/src/simple"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(simple.URL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial %v\n", err)
	}
	defer conn.Close()
	sc := pb.NewSimpleClient(conn)
	client := simple.NewClient(sc)
	if err := client.Greet("Hoge"); err != nil {
		log.Fatalln(err)
	}
}
