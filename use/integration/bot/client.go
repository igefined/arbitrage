package bot

import (
	"context"
	"errors"
)

var (
	ErrSendMessage = errors.New("error send message")
	ErrApiKeyEmpty = errors.New("api key is not exists")
)

type Client interface {
	Send(ctx context.Context, message Message) error
}
