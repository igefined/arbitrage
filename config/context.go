package config

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var ctx context.Context
var once sync.Once

// SigTermIntCtx returns context which cancels by SIGTERM or SIGINT
func SigTermIntCtx() context.Context {
	once.Do(func() {
		var cancel context.CancelFunc

		ctx, cancel = context.WithCancel(context.Background())

		var (
			signalsToIgnore = []os.Signal{syscall.SIGTERM, syscall.SIGINT}
			shutdown        = make(chan os.Signal, len(signalsToIgnore))
		)

		signal.Notify(shutdown, signalsToIgnore...)

		go func() {
			<-shutdown
			cancel()
		}()
	})

	return ctx
}
