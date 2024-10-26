package service

import (
	"app/model"
	"context"
	"time"
)

func Login(ctx context.Context, username, password string) bool {

	time.Sleep(1 * time.Second)

	if username == model.ValidUser.Username && password == model.ValidUser.Password {
		return true
	}
	return false
}