package simple

import (
	"context"
	"fmt"
	pb "github.com/noguchi-hiroshi/grpc-test/proto/simple"
)

type Client interface {
	Greet() error
	Bench()
}

type client struct {
	sc pb.SimpleClient
}

func NewClient(sc pb.SimpleClient) Client {
	return &client{sc: sc}
}

func (c *client) Greet() error {
	ctx := context.Background()
	res, err := c.sc.Greet(ctx, &pb.Request{})
	if err != nil {
		return err
	}
	fmt.Println(res.Message)
	return nil
}

func (c *client) Bench() {
	ctx := context.Background()
	c.sc.Greet(ctx, &pb.Request{})
}
