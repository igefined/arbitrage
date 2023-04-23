package response

import (
	"strconv"

	"github.com/igilgyrg/arbitrage/use/domain"
)

type Ticker struct {
	InstId string `json:"instId"`
	AskPx  string `json:"askPx"`
}

type DailyTicker struct {
	Response
	Data []Ticker `json:"data"`
}

func (t *Ticker) ToDomain() *domain.DailyTicker {
	price, _ := strconv.ParseFloat(t.AskPx, 64)

	return &domain.DailyTicker{
		Symbol: t.InstId,
		Price:  price,
	}
}
