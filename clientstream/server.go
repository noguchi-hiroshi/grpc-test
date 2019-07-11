package clientstream

import (
	"fmt"
	pb "github.com/noguchi-hiroshi/grpc-test/proto/clientstream"
	"io"
)

type service struct {
}

func (s *service) Greet(stream pb.ClientStream_GreetServer) error {
	var sum int32 = 0
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Response{
				Message: fmt.Sprintf("%d", sum),
			})
		}
		if err != nil {
			return err
		}
		sum = sum + res.Num
		fmt.Println(sum)
	}
}

func NewService() pb.ClientStreamServer{
	return &service{}
}
