package repo

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/igilgyrg/arbitrage/schema"
	"github.com/igilgyrg/arbitrage/use/internal/dbo"
)

type BundleRepository interface {
	List(ctx context.Context) ([]dbo.Bundle, error)
	Save(ctx context.Context, bundle *dbo.Bundle) error
}

type repository struct {
	qb *schema.QBuilder
}

func New(qb *schema.QBuilder) BundleRepository {
	return &repository{qb: qb}
}

func (r repository) List(ctx context.Context) (list []dbo.Bundle, err error) {
	if err = pgxscan.Select(ctx, r.qb.Querier(), &list, `select id, symbol, exchange_from, exchange_to, percentage_difference from bundles`); err != nil {
		return
	}

	return
}

func (r repository) Save(ctx context.Context, dbo *dbo.Bundle) (err error) {
	if _, err = r.qb.Querier().Query(ctx, "insert into bundles(symbol, exchange_from, exchange_to, percentage_difference) values($1, $2, $3, $4)", dbo.Symbol, dbo.ExchangeFrom, dbo.ExchangeTo, dbo.PercentageDifference); err != nil {
		return
	}

	return
}
