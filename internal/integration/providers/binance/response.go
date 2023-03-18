package binance

import (
	"strconv"

	"github.com/igilgyrg/arbitrage/internal/domain"
)

type (
	DailyTickerResponse struct {
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

func (t *DailyTickerResponse) ToResponse() *domain.DailyTicker {
	price, _ := strconv.ParseFloat(t.AskPrice, 64)

	return &domain.DailyTicker{
		Symbol: t.Symbol,
		Price:  price,
	}
}
