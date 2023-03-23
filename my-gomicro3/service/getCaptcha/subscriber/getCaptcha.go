package subscriber

import (
	"context"

	log "github.com/asim/go-micro/v3/logger"

	getCaptcha "my-gomicro3/service/get-captcha/proto/getCaptcha"
)

type GetCaptcha struct{}

func (e *GetCaptcha) Handle(ctx context.Context, msg *getCaptcha.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *getCaptcha.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
