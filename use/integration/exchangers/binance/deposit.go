package binance

import (
	"context"
	"time"
)

func (c *client) DepositNetwork(ctx context.Context, symbol string) (networks []string) {
	if len(c.allCoinsInfo) == 0 || time.Now().Hour() == 0 {
		err := c.AllCoinsInfoProcessed(ctx)
		if err != nil {
			c.logger.Error(err)

			return
		}
	}

	info := c.allCoinsInfo[symbol]
	if len(info.NetworkList) == 0 {
		c.logger.Warnf("binance: symbol %s does not support any network", symbol)

		return
	}

	for _, n := range info.NetworkList {
		if n.DepositEnable {
			networks = append(networks, n.Name)
		}
	}

	return
}
