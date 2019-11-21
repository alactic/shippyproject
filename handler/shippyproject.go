package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	shippyproject "github.com/alactic/shippyproject/proto/shippyproject"
)

type Shippyproject struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Shippyproject) Call(ctx context.Context, req *shippyproject.Request, rsp *shippyproject.Response) error {
	log.Log("Received Shippyproject.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Shippyproject) Stream(ctx context.Context, req *shippyproject.StreamingRequest, stream shippyproject.Shippyproject_StreamStream) error {
	log.Logf("Received Shippyproject.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&shippyproject.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Shippyproject) PingPong(ctx context.Context, stream shippyproject.Shippyproject_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&shippyproject.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
