package main

import (
	pb "github.com/noguchi-hiroshi/grpc-test/proto/serverstream"
	"github.com/noguchi-hiroshi/grpc-test/serverstream"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", serverstream.URL)
	if err != nil {
		log.Fatalf("failed to start %v", err)
	}
	server := grpc.NewServer()
	srv := serverstream.NewService()
	pb.RegisterServerStreamServer(server, srv)
	log.Println("start")
	if err := server.Serve(lis); err != nil {
		log.Fatalln("failed to server")
	}
}
