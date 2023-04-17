package kucoin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitSymbol(t *testing.T) {
	tCases := []struct {
		symbol   string
		expected string
	}{
		{
			symbol:   "BTCUSDT",
			expected: "BTC-USDT",
		},
		{
			symbol:   "BTCUSD",
			expected: "BTC-USD",
		},
		{
			symbol:   "BTCBUSD",
			expected: "BTC-BUSD",
		},
		{
			symbol:   "ETHUSDT",
			expected: "ETH-USDT",
		}, {
			symbol:   "ETHUSD",
			expected: "ETH-USD",
		},
		{
			symbol:   "ETHBUSD",
			expected: "ETH-BUSD",
		},
		{
			symbol:   "ARBBUSD",
			expected: "ARB-BUSD",
		},
	}

	for _, c := range tCases {
		assert.Equal(t, c.expected, SplitSymbol(c.symbol))
	}
}
