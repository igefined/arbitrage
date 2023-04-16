package main

import (
	"time"

	"github.com/igilgyrg/arbitrage/api"
	"github.com/igilgyrg/arbitrage/api/endpoints"
	"github.com/igilgyrg/arbitrage/config"
	"github.com/igilgyrg/arbitrage/log"
	"github.com/igilgyrg/arbitrage/schema"
	"github.com/igilgyrg/arbitrage/use"
	"github.com/igilgyrg/arbitrage/use/integration/ninja"
	"github.com/igilgyrg/arbitrage/use/service/inspector"
	"github.com/igilgyrg/arbitrage/use/service/scheduler"

	"go.uber.org/fx"
)

func main() {
	ctx := config.SigTermIntCtx()

	logger := log.New()
	cfg := config.New()

	app := fx.New(
		fx.Supply(logger, cfg, inspector.DefaultExchangers(logger)),

		fx.Provide(
			schema.New,
			api.NewServer,
			endpoints.New,
			ninja.New,
			use.NewComposite,
		),

		use.Constructor(),

		fx.Invoke(func(_ *api.Server, qb *schema.QBuilder, scheduler scheduler.Service) {
			schema.Migrate(logger, &schema.DB, qb.ConnString())

			scheduler.TemporalArbitrage(ctx, time.Minute*5)
			scheduler.TemporalSymbols(ctx, time.Hour*24)
		}),
	)

	app.Run()
}
