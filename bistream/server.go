package bistream

import (
	pb "github.com/noguchi-hiroshi/grpc-test/proto/bistream"
	"io"
)

type service struct {
}

func (s *service) Greet(stream pb.BiStream_GreetServer) error {
	for {
		res, err := stream.Recv()
		// Notice: err != nil より先に判定しないと, err を返してしまうので注意
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		if res.Finished == true {
			return nil
		}

		err = stream.Send(&pb.Response{
			Message: "それな > " + res.Message,
		})

		if err != nil {
			return err
		}
	}
}

func NewService() pb.BiStreamServer {
	return &service{}
}
