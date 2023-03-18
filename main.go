package main

import (
	"github.com/igilgyrg/arbitrage/api"
	"github.com/igilgyrg/arbitrage/api/endpoints"
	"github.com/igilgyrg/arbitrage/config"
	"github.com/igilgyrg/arbitrage/log"
	"github.com/igilgyrg/arbitrage/schema"
	"github.com/igilgyrg/arbitrage/use"
	"github.com/igilgyrg/arbitrage/use/service/inspector"
	"go.uber.org/fx"
)

func main() {
	logger := log.New()
	cfg := config.New()

	app := fx.New(
		fx.Supply(logger, cfg, inspector.DefaultExchangers(logger)),

		fx.Provide(
			schema.New,
			api.NewServer,
			endpoints.New,
			use.NewComposite,
		),

		use.Constructor(),

		fx.Invoke(func(_ *api.Server, qb *schema.QBuilder) {
			schema.Migrate(logger, &schema.DB, qb.ConnString())
		}),
	)

	app.Run()
}
