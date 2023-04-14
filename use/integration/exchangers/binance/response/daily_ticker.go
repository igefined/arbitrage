package response

import (
	"strconv"

	"github.com/igilgyrg/arbitrage/use/domain"
)

type (
	DailyTicker struct {
		Symbol             string `json:"symbol"`
		PriceChange        string `json:"priceChange"`
		PriceChangePercent string `json:"priceChangePercent"`
		LastPrice          string `json:"lastPrice"`
		AskPrice           string `json:"askPrice"`
	}

	ErrorResponse struct {
		Code    int    `json:"code"`
		Message string `json:"msg"`
	}
)

func (t *DailyTicker) ToDomain() *domain.DailyTicker {
	price, _ := strconv.ParseFloat(t.AskPrice, 64)

	return &domain.DailyTicker{
		Symbol: t.Symbol,
		Price:  price,
	}
}
