package usymbol

import (
	"fmt"
	"strings"
)

var availableToken = []string{"USDT", "BUSD", "USD"}

func SplitSymbol(symbol string) string {
	for _, t := range availableToken {
		ok := strings.HasSuffix(symbol, t)
		if ok {
			return fmt.Sprintf("%s-%s", strings.Split(symbol, t)[0], t)
		}
	}

	return symbol
}
