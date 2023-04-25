package usymbol

import (
	"fmt"
	"strings"
)

var availableToken = []string{"USDT", "BUSD", "USD", "USDC"}

func SeparateSymbol(symbol, separator string) string {
	for _, t := range availableToken {
		ok := strings.HasSuffix(symbol, t)
		if ok {
			return fmt.Sprintf("%s%s%s", strings.Split(symbol, t)[0], separator, t)
		}
	}

	return symbol
}

func Split(symbol string) (crypto string, stable string) {
	for _, t := range availableToken {
		ok := strings.HasSuffix(symbol, t)
		if ok {
			crypto = strings.Split(symbol, t)[0]
			stable = strings.Split(symbol, t)[0]

			return
		}
	}

	return
}
