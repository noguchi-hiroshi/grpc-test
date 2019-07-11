package main

import (
	pb "github.com/noguchi-hiroshi/grpc-test/proto/clientstream"
	"github.com/noguchi-hiroshi/grpc-test/clientstream"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", clientstream.URL)
	if err != nil {
		log.Fatalf("failed to start %v", err)
	}
	server := grpc.NewServer()
	srv := clientstream.NewService()
	pb.RegisterClientStreamServer(server, srv)
	log.Println("start")
	if err := server.Serve(lis); err != nil {
		log.Fatalln("failed to server")
	}
}
