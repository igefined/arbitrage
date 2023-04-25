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

	Clear(ctx context.Context) error
}

type repository struct {
	qb *schema.QBuilder
}

func New(qb *schema.QBuilder) BundleRepository {
	return &repository{qb: qb}
}

func (r *repository) List(ctx context.Context) (list []dbo.Bundle, err error) {
	q := `select id, symbol, exchange_from, price_from, exchange_to, price_to, percentage_difference, updated_at, deposit_networks, withdraw_networks from bundles`
	if err = pgxscan.Select(ctx, r.qb.Querier(), &list, q); err != nil {
		return
	}

	return
}

func (r *repository) Save(ctx context.Context, dbo *dbo.Bundle) error {
	q := `insert into bundles(symbol, exchange_from, price_from, exchange_to, price_to, percentage_difference, deposit_networks, withdraw_networks) values($1, $2, $3, $4, $5, $6, $7, $8)`
	rows, err := r.qb.Querier().Query(ctx, q, dbo.Symbol, dbo.ExchangeFrom, dbo.PriceFrom, dbo.ExchangeTo, dbo.PriceTo, dbo.PercentageDifference, dbo.DepositNetworks, dbo.WithdrawNetworks)
	if err != nil {
		return err
	}
	defer rows.Close()

	return err
}

func (r *repository) Clear(ctx context.Context) error {
	rows, err := r.qb.Querier().Query(ctx, `delete from bundles where id > 0`)
	if err != nil {
		return err
	}
	defer rows.Close()

	return err
}
