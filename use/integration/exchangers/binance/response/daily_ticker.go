package response

import (
	"strconv"

	"github.com/igilgyrg/arbitrage/use/domain"
)

type DailyTicker struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func (t *DailyTicker) ToDomain() *domain.DailyTicker {
	price, _ := strconv.ParseFloat(t.Price, 64)

	return &domain.DailyTicker{
		Symbol: t.Symbol,
		Price:  price,
	}
}
