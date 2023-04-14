package symbol

import (
	"context"
	"strings"
)

func (c *client) Upgrade(ctx context.Context) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	symbols, err := c.ninja.CryptoSymbols(ctx)
	if err != nil {
		c.log.Errorf("symbols: err get symbsol: %v", err)

		return
	}

	if len(symbols) == 0 {
		return
	}

	tempSymbols := make([]string, 0, len(symbols))
	for _, symbol := range symbols {
		if validate(symbol) {
			tempSymbols = append(tempSymbols, symbol)
		}
	}

	c.symbols = tempSymbols
}

func validate(symbol string) (out bool) {
	out = strings.HasSuffix(symbol, "USD")
	out = out || strings.HasSuffix(symbol, "USDT")
	out = out || strings.HasSuffix(symbol, "USDC")
	out = out || strings.HasSuffix(symbol, "BUSD")

	return
}
