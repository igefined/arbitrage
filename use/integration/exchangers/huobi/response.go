package mexc

import (
	"strconv"

	"github.com/igilgyrg/arbitrage/use/domain"
)

type (
	DailyTickerResponse struct {
		Symbol   string `json:"symbol"`
		BidPrice string `json:"bidPrice"`
		BidQty   string `json:"bidQty"`
		AskPrice string `json:"askPrice"`
		AskQty   string `json:"askQty"`
	}

	ErrorResponse struct {
		Code    int    `json:"code"`
		Message string `json:"msg"`
	}
)

func (t *DailyTickerResponse) ToResponse() *domain.DailyTicker {
	price, _ := strconv.ParseFloat(t.AskPrice, 64)

	return &domain.DailyTicker{
		Symbol: t.Symbol,
		Price:  price,
	}
}
