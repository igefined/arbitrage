package main

import (
	"github.com/igilgyrg/arbitrage/api"
	"github.com/igilgyrg/arbitrage/api/endpoints"
	"github.com/igilgyrg/arbitrage/config"
	"github.com/igilgyrg/arbitrage/log"

	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/fx"
)

func main() {
	logger := log.New()
	cfg := config.New()

	app := fx.New(
		fx.Supply(logger, cfg),

		fx.Provide(
			api.NewServer,
			endpoints.New,
		),

		fx.Invoke(func(_ *api.Server) {}),
	)

	app.Run()
}
