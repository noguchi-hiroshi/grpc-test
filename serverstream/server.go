package serverstream

import (
	"errors"
	"fmt"
	pb "github.com/noguchi-hiroshi/grpc-test/proto/serverstream"
)

type service struct {
}

func (s *service) Greet(req *pb.Request, stream pb.ServerStream_GreetServer) error {
	if req.Min > req.Max {
		return errors.New("bad min is greater than max")
	}
	for i := req.Min; i < req.Max; i++ {
		res := pb.Response{Message: fmt.Sprintf("number: %d", i)}
		if err := stream.Send(&res); err != nil {
			return err
		}
	}
	return nil
}

func NewService() pb.ServerStreamServer {
	return &service{}
}
