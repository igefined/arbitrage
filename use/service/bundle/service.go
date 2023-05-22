package bundle

import (
	"context"

	"github.com/igdotog/core/logger"
	"github.com/igilgyrg/arbitrage/use/domain"
	"github.com/igilgyrg/arbitrage/use/internal/repo"
)

type Service interface {
	List(ctx context.Context) ([]domain.Bundle, error)
	Save(ctx context.Context, bundle *domain.Bundle) error

	Clear(ctx context.Context) error
}

type service struct {
	logger *logger.Logger
	bundle repo.BundleRepository
}

func New(logger *logger.Logger, bundle repo.BundleRepository) Service {
	return &service{logger: logger, bundle: bundle}
}
