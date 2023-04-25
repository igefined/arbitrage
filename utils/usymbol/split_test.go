package usymbol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeparateSymbol(t *testing.T) {
	tCases := []struct {
		symbol    string
		separator string
		expected  string
	}{
		{
			symbol:    "BTCUSDT",
			separator: "-",
			expected:  "BTC-USDT",
		},
		{
			symbol:    "BTCUSD",
			separator: "-",
			expected:  "BTC-USD",
		},
		{
			symbol:    "BTCBUSD",
			separator: "-",
			expected:  "BTC-BUSD",
		},
		{
			symbol:    "ETHUSDT",
			separator: "-",
			expected:  "ETH-USDT",
		}, {
			symbol:    "ETHUSD",
			separator: "-",
			expected:  "ETH-USD",
		},
		{
			symbol:    "ETHBUSD",
			separator: "-",
			expected:  "ETH-BUSD",
		},
		{
			symbol:    "ARBBUSD",
			separator: "-",
			expected:  "ARB-BUSD",
		},
	}

	for _, c := range tCases {
		assert.Equal(t, c.expected, SeparateSymbol(c.symbol, c.separator))
	}
}
