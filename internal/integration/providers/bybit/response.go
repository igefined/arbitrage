package bybit

import (
	"strconv"

	"github.com/igilgyrg/arbitrage/internal/domain"
)

type (
	Response struct {
		Code          int         `json:"retCode"`
		Message       string      `json:"retMsg"`
		Result        interface{} `json:"result"`
		ExtensionInfo interface{} `json:"retExtInfo"`
		Time          int64       `json:"time"`
	}

	TickerResponse struct {
		Category string   `json:"category"`
		List     []Ticker `json:"list"`
	}

	Ticker struct {
		Symbol   string `json:"symbol"`
		AskPrice string `json:"ask1Price"`
	}
)

func (t *Ticker) ToResponse() *domain.DailyTicker {
	price, _ := strconv.ParseFloat(t.AskPrice, 64)

	return &domain.DailyTicker{
		Symbol: t.Symbol,
		Price:  price,
	}
}
