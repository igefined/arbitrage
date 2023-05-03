package huobi

import (
	"context"
	"time"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/igilgyrg/arbitrage/log"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers"

	huobi "github.com/huobirdcenter/huobi_golang/pkg/client"
)

const (
	ExchangeName = "huobi"
	host         = "api.huobi.pro"
)

type (
	client struct {
		account *huobi.AccountClient
		market  *huobi.MarketClient
		wallet  *huobi.WalletClient
		cfg     *config
		logger  *log.Logger
	}

	config struct {
		ApiKey    string `config:"HUOBI_API_KEY"`
		SecretKey string `config:"HUOBI_SECRET_KEY"`
	}
)

func New(logger *log.Logger) exchangers.Client {
	cfg := &config{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := confita.NewLoader(env.NewBackend()).Load(ctx, cfg); err != nil {
		logger.Error(err)
	}

	accountClient := new(huobi.AccountClient).Init(cfg.ApiKey, cfg.SecretKey, host)
	marketClient := new(huobi.MarketClient).Init(host)
	walletClient := new(huobi.WalletClient).Init(cfg.ApiKey, cfg.SecretKey, host)

	return &client{account: accountClient, market: marketClient, wallet: walletClient, cfg: cfg, logger: logger}
}

func (c *client) Name() string {
	return ExchangeName
}
