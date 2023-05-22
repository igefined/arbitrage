package domain

import (
	"fmt"
	"time"
)

type Bundle struct {
	Id                   int       `json:"id"`
	Symbol               string    `json:"symbol"`
	ExchangeFrom         string    `json:"exchange_from"`
	PriceFrom            float64   `json:"price_from"`
	ExchangeTo           string    `json:"exchange_to"`
	PriceTo              float64   `json:"price_to"`
	PercentageDifference float64   `json:"percentage_difference"`
	UpdatedAt            time.Time `json:"updated_at"`
	DepositNetworks      []string  `json:"deposit_networks"`
	WithdrawNetworks     []string  `json:"withdraw_networks"`
}

func (b *Bundle) String() string {
	return fmt.Sprintf("SYMBOL: %s\nFROM(%s)-TO(%s): %s-%s\nPRICES: %f-%f\nPERCENT: %f", b.Symbol, b.ExchangeFrom, b.ExchangeTo, b.DepositNetworks, b.WithdrawNetworks, b.PriceFrom, b.PriceTo, b.PercentageDifference)
}
