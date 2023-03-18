package dbo

type Bundle struct {
	Id                   int     `db:"id"`
	Symbol               string  `db:"symbol"`
	ExchangeFrom         string  `db:"exchange_from"`
	ExchangeTo           string  `db:"exchange_to"`
	PercentageDifference float64 `db:"percentage_difference"`
}
