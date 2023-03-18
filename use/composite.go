package use

import (
	"github.com/igilgyrg/arbitrage/use/internal/repo"
	"github.com/igilgyrg/arbitrage/use/service/bundle"
	"github.com/igilgyrg/arbitrage/use/service/inspector"
	"go.uber.org/fx"
)

func Constructor() fx.Option {
	return fx.Provide(
		repo.New,
		bundle.New,
		inspector.New,
	)
}

type UseCase interface {
	Bundles() bundle.Service
	Inspector() inspector.Service
}

type useCase struct {
	bundles   bundle.Service
	inspector inspector.Service
}

func NewComposite(
	bundles bundle.Service,
	inspector inspector.Service,
) UseCase {
	return &useCase{
		bundles:   bundles,
		inspector: inspector,
	}
}

func (u *useCase) Bundles() bundle.Service {
	return u.bundles
}

func (u *useCase) Inspector() inspector.Service {
	return u.inspector
}
