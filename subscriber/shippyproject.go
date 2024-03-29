package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	shippyproject "github.com/alactic/shippyproject/proto/shippyproject"
)

type Shippyproject struct{}

func (e *Shippyproject) Handle(ctx context.Context, msg *shippyproject.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *shippyproject.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
