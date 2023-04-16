package dbo

import "time"

type Bundle struct {
	Id                   int       `db:"id"`
	Symbol               string    `db:"symbol"`
	ExchangeFrom         string    `db:"exchange_from"`
	PriceFrom            float64   `db:"price_from"`
	ExchangeTo           string    `db:"exchange_to"`
	PriceTo              float64   `db:"price_to"`
	PercentageDifference float64   `db:"percentage_difference"`
	UpdatedAt            time.Time `db:"updated_at"`
}
