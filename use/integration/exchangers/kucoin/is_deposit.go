package kucoin

import "context"

func (c *client) IsDeposit(ctx context.Context, symbol string) bool {
	return true
}
