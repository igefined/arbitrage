package huobi

import (
	"context"
	"strings"

	"github.com/igilgyrg/arbitrage/use/integration/exchangers"
)

func (c *client) WithdrawNetwork(ctx context.Context, symbol string) (networks []string) {
	resp, err := c.wallet.GetWithdrawQuota(strings.ToLower(symbol))
	if err != nil {
		c.logger.Errorf("huobi deposit network request: %v", err)

		return
	}

	if resp == nil {
		c.logger.Error("huobi deposit network: nil result")

		return
	}

	if len(resp.Chains) == 0 {
		c.logger.Errorf("huobi deposit network: %v", exchangers.ErrSymbolNotFound)

		return
	}

	networks = make([]string, 0, len(resp.Chains))
	for i := range resp.Chains {
		chain := resp.Chains[i]
		networks = append(networks, chain.Chain)
	}

	return
}
