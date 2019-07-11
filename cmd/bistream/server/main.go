package main

import (
	"github.com/noguchi-hiroshi/grpc-test/bistream"
	pb "github.com/noguchi-hiroshi/grpc-test/proto/bistream"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", bistream.URL)
	if err != nil {
		log.Fatalf("failed to start %v", err)
	}
	server := grpc.NewServer()
	srv := bistream.NewService()
	pb.RegisterBiStreamServer(server, srv)
	log.Println("start")
	if err := server.Serve(lis); err != nil {
		log.Fatalln("failed to server")
	}
}
