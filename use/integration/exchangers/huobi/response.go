package huobi

import (
	"github.com/igilgyrg/arbitrage/use/domain"
)

type (
	Response struct {
		Ch         string      `json:"ch"`
		Code       int         `json:"ts"`
		Status     string      `json:"status"`
		Result     interface{} `json:"tick"`
		ErrMessage string      `json:"err-msg"`
		ErrCode    string      `json:"err-code"`
	}

	DailyTickerResponse struct {
		Id  int       `json:"id"`
		Ask []float64 `json:"ask"`
	}
)

func (t *DailyTickerResponse) ToResponse(symbol string) *domain.DailyTicker {
	var price float64
	if len(t.Ask) > 0 {
		price = t.Ask[0]
	}

	return &domain.DailyTicker{
		Symbol: symbol,
		Price:  price,
	}
}
