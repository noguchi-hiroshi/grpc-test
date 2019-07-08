package simple

import (
	"context"
	"fmt"
	pb "github.com/noguchi-hiroshi/grpc-test/proto/simple"
)

type server struct {
}

func (server) Greet(_ context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{
		Message: fmt.Sprintf("Hello %s", req.Nickname),
	}, nil
}

func NewService() pb.SimpleServer {
	return &server{}
}
