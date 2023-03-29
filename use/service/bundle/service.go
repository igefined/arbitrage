package bundle

import (
	"context"

	"github.com/igilgyrg/arbitrage/log"
	"github.com/igilgyrg/arbitrage/use/domain"
	"github.com/igilgyrg/arbitrage/use/internal/repo"
)

type Service interface {
	List(ctx context.Context) ([]domain.Bundle, error)
	Save(ctx context.Context, bundle *domain.Bundle) error

	Clear(ctx context.Context) error
}

type service struct {
	logger *log.Logger
	bundle repo.BundleRepository
}

func New(logger *log.Logger, bundle repo.BundleRepository) Service {
	return &service{logger: logger, bundle: bundle}
}
