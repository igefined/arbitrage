package response

import (
	"strconv"

	"github.com/igilgyrg/arbitrage/use/domain"
)

type (
	TickerResponse struct {
		Category string   `json:"category"`
		List     []Ticker `json:"list"`
	}

	Ticker struct {
		Symbol   string `json:"symbol"`
		AskPrice string `json:"ask1Price"`
	}
)

func (t *Ticker) ToDomain() *domain.DailyTicker {
	price, _ := strconv.ParseFloat(t.AskPrice, 64)

	return &domain.DailyTicker{
		Symbol: t.Symbol,
		Price:  price,
	}
}
