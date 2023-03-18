package domain

type Bundle struct {
	Id                   int     `json:"id"`
	Symbol               string  `json:"symbol"`
	ExchangeFrom         string  `json:"exchange_from"`
	ExchangeTo           string  `json:"exchange_to"`
	PercentageDifference float64 `json:"percentage_difference"`
}
