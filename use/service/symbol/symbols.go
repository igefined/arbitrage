package symbol

import "context"

func (c *client) Symbols(ctx context.Context) []string {
	return c.symbols
}
