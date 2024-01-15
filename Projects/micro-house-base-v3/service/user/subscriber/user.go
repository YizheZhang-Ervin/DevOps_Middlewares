package subscriber

import (
	"context"

	log "github.com/asim/go-micro/v3/logger"

	user "micro-house-base-v3/service/user/proto/user"
)

type User struct{}

func (e *User) Handle(ctx context.Context, msg *user.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *user.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
