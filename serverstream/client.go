package serverstream

import (
	"context"
	pb "github.com/noguchi-hiroshi/grpc-test/proto/serverstream"
	"io"
	"log"
	"time"
)

type Client interface {
	Greet(min int32, max int32) error
}

type client struct {
	sc pb.ServerStreamClient
}

func NewClient(sc pb.ServerStreamClient) Client {
	return &client{sc: sc}
}

func (c *client) Greet(min int32, max int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.sc.Greet(ctx, &pb.Request{
		Min: min,
		Max: max,
	})
	if err != nil {
		return err
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf(res.Message)
	}
	return nil
}
