package main

import (
	"time"

	"github.com/igdotog/core/config"
	"github.com/igdotog/core/logger"
	"github.com/igilgyrg/arbitrage/api"
	"github.com/igilgyrg/arbitrage/api/endpoints"
	"github.com/igilgyrg/arbitrage/schema"
	"github.com/igilgyrg/arbitrage/use"
	"github.com/igilgyrg/arbitrage/use/integration/bot/telegram"
	"github.com/igilgyrg/arbitrage/use/integration/ninja"
	"github.com/igilgyrg/arbitrage/use/service/inspector"
	"github.com/igilgyrg/arbitrage/use/service/scheduler"

	"go.uber.org/fx"
)

func main() {
	ctx := config.SigTermIntCtx()

	log := logger.New()
	cfg := config.NewBaseConfig(ctx)

	app := fx.New(
		fx.Supply(log, cfg, inspector.DefaultExchangers(log)),

		fx.Provide(
			schema.New,
			api.NewServer,
			endpoints.New,
			ninja.New,
			use.NewComposite,
			telegram.New,
		),

		use.Constructor(),

		fx.Invoke(func(_ *api.Server, qb *schema.QBuilder, scheduler scheduler.Service) {
			schema.Migrate(log, &schema.DB, qb.ConnString())

			scheduler.TemporalArbitrage(ctx, time.Minute*1)
			scheduler.TemporalSymbols(ctx, time.Hour*24)
		}),
	)

	app.Run()
}
