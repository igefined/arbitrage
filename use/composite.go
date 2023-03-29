package use

import (
	"github.com/igilgyrg/arbitrage/use/internal/repo"
	"github.com/igilgyrg/arbitrage/use/service/bundle"
	"github.com/igilgyrg/arbitrage/use/service/inspector"
	"github.com/igilgyrg/arbitrage/use/service/scheduler"

	"go.uber.org/fx"
)

func Constructor() fx.Option {
	return fx.Provide(
		repo.New,
		bundle.New,
		inspector.New,
		scheduler.New,
	)
}

type UseCase interface {
	Bundles() bundle.Service
	Inspector() inspector.Service
	Scheduler() scheduler.Service
}

type useCase struct {
	bundles   bundle.Service
	inspector inspector.Service
	scheduler scheduler.Service
}

func NewComposite(
	bundles bundle.Service,
	inspector inspector.Service,
	scheduler scheduler.Service,
) UseCase {
	return &useCase{
		bundles:   bundles,
		inspector: inspector,
		scheduler: scheduler,
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
