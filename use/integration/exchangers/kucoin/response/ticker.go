package response

import (
	"strconv"

	"github.com/igilgyrg/arbitrage/use/domain"
)

type DailyTicker struct {
	Symbol string `json:"symbol"`
	Time   int64  `json:"time"`
	Buy    string `json:"buy"`
	Sell   string `json:"sell"`
}

func (t *DailyTicker) ToDomain() *domain.DailyTicker {
	price, _ := strconv.ParseFloat(t.Sell, 64)

	return &domain.DailyTicker{
		Symbol: t.Symbol,
		Price:  price,
	}
}
