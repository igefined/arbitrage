//go:build units

package symbol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	tCases := []struct {
		symbol   string
		expected bool
	}{
		{
			symbol:   "BTCUSDT",
			expected: true,
		},
		{
			symbol:   "BTCUSDC",
			expected: false,
		},
		{
			symbol:   "BTCBUSDT",
			expected: true,
		},

		{
			symbol:   "BTCBUSD4",
			expected: false,
		},
	}

	for _, c := range tCases {
		assert.Equal(t, c.expected, validate(c.symbol))
	}
}
