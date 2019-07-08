package main

import (
	pb "github.com/noguchi-hiroshi/grpc-test/proto/simple"
	"github.com/noguchi-hiroshi/grpc-test/src/simple"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", simple.URL)
	if err != nil {
		log.Fatalf("failed to start %v", err)
	}
	server := grpc.NewServer()
	srv := simple.NewService()
	pb.RegisterSimpleServer(server, srv)
	log.Println("start")
	if err := server.Serve(lis); err != nil {
		log.Fatalln("failed to server")
	}
}
