package bundle

import (
	"context"
	"fmt"

	"github.com/igilgyrg/arbitrage/use/domain"
)

func (s *service) List(ctx context.Context) ([]domain.Bundle, error) {
	list, err := s.bundle.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("bundle service: %w", err)
	}

	result := make([]domain.Bundle, len(list))
	for i, b := range list {
		result[i] = domain.Bundle{
			Id:                   b.Id,
			Symbol:               b.Symbol,
			ExchangeFrom:         b.ExchangeFrom,
			ExchangeTo:           b.ExchangeTo,
			PriceFrom:            b.PriceFrom,
			PriceTo:              b.PriceTo,
			PercentageDifference: b.PercentageDifference,
			UpdatedAt:            b.UpdatedAt,
		}
	}

	return result, nil
}
