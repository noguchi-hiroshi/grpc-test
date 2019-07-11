package bistream

import (
	"context"
	"fmt"
	pb "github.com/noguchi-hiroshi/grpc-test/proto/bistream"
	"io"
	"log"
	"time"
)

type Client interface {
	Greet() error
}

type client struct {
	sc pb.BiStreamClient
}

func NewClient(sc pb.BiStreamClient) Client {
	return &client{sc: sc}
}

func (c *client) Greet() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)

	defer cancel()

	stream, err := c.sc.Greet(ctx)

	if err != nil {
		return err
	}

	// 値の受け渡しが不要な場合のchannelには空structがよく使われます。サイズがゼロだからです。
	done := make(chan struct{}, 0)

	// 受信は並行処理で行う
	go c.recv(stream, done)

	// server に close を通知したが失敗した場合は強制的に close させる
	if err = c.send(stream); err != nil {
		close(done)
	}

	<- done

	return err
}

func (c *client) recv(stream pb.BiStream_GreetClient, done chan struct{}) {
	defer close(done)
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			log.Println("closed")
			return
		}
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("Server: %s \n", res.Message)
	}
}

func (c *client) send(stream pb.BiStream_GreetClient) error {
	for i := 1; i <= 10; i++ {
		mes := fmt.Sprintf("進捗どうですか？ (%d回目)", i)
		fmt.Println("Client: " + mes)
		err := stream.Send(&pb.Request{
			Message: mes,
			Finished: false,
		})
		if err != nil {
			break
		}
		time.Sleep(time.Millisecond * 200)
	}

	return stream.CloseSend()
}

