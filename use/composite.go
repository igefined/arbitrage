package use

import (
	"github.com/igilgyrg/arbitrage/use/internal/repo"
	"github.com/igilgyrg/arbitrage/use/service/bundle"
	"github.com/igilgyrg/arbitrage/use/service/inspector"
	"github.com/igilgyrg/arbitrage/use/service/scheduler"
	"github.com/igilgyrg/arbitrage/use/service/symbol"

	"go.uber.org/fx"
)

func Constructor() fx.Option {
	return fx.Provide(
		repo.New,
		bundle.New,
		inspector.New,
		scheduler.New,
		symbol.New,
	)
}

type UseCase interface {
	Bundles() bundle.Service
	Inspector() inspector.Service
	Scheduler() scheduler.Service
	Symbols() symbol.Service
}

type useCase struct {
	bundles   bundle.Service
	inspector inspector.Service
	scheduler scheduler.Service
	symbols   symbol.Service
}

func NewComposite(
	bundles bundle.Service,
	inspector inspector.Service,
	scheduler scheduler.Service,
	symbols symbol.Service,
) UseCase {
	return &useCase{
		bundles:   bundles,
		inspector: inspector,
		scheduler: scheduler,
		symbols:   symbols,
	}
}

func (u *useCase) Bundles() bundle.Service {
	return u.bundles
}

func (u *useCase) Inspector() inspector.Service {
	return u.inspector
}

func (u *useCase) Scheduler() scheduler.Service {
	return u.scheduler
}

func (u *useCase) Symbols() symbol.Service {
	return u.symbols
}
