package domain

type (
	Spreads struct {
		Symbol  string
		Spreads []SpreadInfo
	}

	SpreadInfo struct {
		ExchangeName string
		Price        float64
	}
)
