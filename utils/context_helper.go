package utils

import (
	"context"
	"time"
)

func InitContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	return ctx, cancel
}
