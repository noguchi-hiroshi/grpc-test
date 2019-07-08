package simple

import (
	"context"
	"fmt"
	pb "github.com/noguchi-hiroshi/grpc-test/proto/simple"
)

type Client interface {
	Greet(nickname string) error
}

type client struct {
	sc pb.SimpleClient
}

func NewClient(sc pb.SimpleClient) Client {
	return &client{sc: sc}
}

func (c *client) Greet(nickname string) error {
	ctx := context.Background()
	res, err := c.sc.Greet(ctx, &pb.Request{Nickname: nickname})
	if err != nil {
		return err
	}
	fmt.Println(res.Message)
	return nil
}
