package clientstream

import (
	"context"
	pb "github.com/noguchi-hiroshi/grpc-test/proto/clientstream"
	"log"
	"time"
)

type Client interface {
	Greet() error
}

type client struct {
	sc pb.ClientStreamClient
}

func NewClient(sc pb.ClientStreamClient) Client {
	return &client{sc: sc}
}

func (c *client) Greet() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.sc.Greet(ctx)
	if err != nil {
		return err
	}

	var i int32
	for i = 1; i <= 5; i++ {
		if err := stream.Send(&pb.Request{Num: i}); err != nil {
			return err
		}
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		return err
	}

	log.Println(res.Message)

	return nil
}
