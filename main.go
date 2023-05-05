package main

import (
	"fmt"

	"github.com/igilgyrg/arbitrage/config"
	"github.com/igilgyrg/arbitrage/log"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers/kucoin"
)

func main() {
	ctx := config.SigTermIntCtx()

	logger := log.New()
	//cfg := config.New()

	network := kucoin.New(logger).WithdrawNetwork(ctx, "TON")
	for i := range network {
		fmt.Println(network[i])
	}

	//app := fx.New(
	//	fx.Supply(logger, cfg, inspector.DefaultExchangers(logger)),
	//
	//	fx.Provide(
	//		schema.New,
	//		api.NewServer,
	//		endpoints.New,
	//		ninja.New,
	//		use.NewComposite,
	//		telegram.New,
	//	),
	//
	//	use.Constructor(),
	//
	//	fx.Invoke(func(_ *api.Server, qb *schema.QBuilder, scheduler scheduler.Service) {
	//		schema.Migrate(logger, &schema.DB, qb.ConnString())
	//
	//		scheduler.TemporalArbitrage(ctx, time.Minute*1)
	//		scheduler.TemporalSymbols(ctx, time.Hour*24)
	//	}),
	//)
	//
	//app.Run()
}
