//go:build units

package schema

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/igdotog/core/config"
	"github.com/igdotog/core/logger"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
)

func TestGetDatabaseName(t *testing.T) {
	const (
		url    = "postgres://postgres:postgres@localhost:5466/test_clients?sslmode=disable"
		expect = "test_clients"
	)

	assert.Equal(t, GetDatabaseName(url), expect)
}

func TestReplaceDbName(t *testing.T) {
	tCases := []struct {
		srcUrl    string
		dbName    string
		resultUrl string
	}{
		{
			srcUrl:    "postgres://postgres:postgres@localhost:5466/test?sslmode=disable",
			dbName:    "silly",
			resultUrl: "postgres://postgres:postgres@localhost:5466/silly?sslmode=disable",
		},
		{
			srcUrl:    "postgres://postgres:postgres@localhost:5466/test",
			dbName:    "silly",
			resultUrl: "postgres://postgres:postgres@localhost:5466/silly",
		},
		{
			srcUrl:    "postgres://postgres:12345@localhost:5432/common?sslmode=disable&pool_max_conns=16&pool_max_conn_idle_time=30m&pool_max_conn_lifetime=1h&pool_health_check_period=1m",
			dbName:    "nh_common",
			resultUrl: "postgres://postgres:12345@localhost:5432/nh_common?sslmode=disable&pool_max_conns=16&pool_max_conn_idle_time=30m&pool_max_conn_lifetime=1h&pool_health_check_period=1m",
		},
	}

	for _, c := range tCases {
		assert.Equal(t, ReplaceDbName(c.srcUrl, c.dbName), c.resultUrl)
	}
}

func TestCreateDatabase(t *testing.T) {
	var (
		isExists bool
		logger   = logger.New()
	)

	cfg := config.NewBaseConfig(context.Background())
	url := cfg.DBUrl
	assert.NotEmpty(t, url)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	CreateDatabase(ctx, logger, url)

	pool, err := pgxpool.New(ctx, ReplaceDbName(url, "postgres"))
	assert.NoError(t, err)

	checkingSql := `select exists(select datname from pg_catalog.pg_database where datname = $1) as exist`
	row := pool.QueryRow(ctx, checkingSql, GetDatabaseName(url))
	err = row.Scan(&isExists)
	assert.NoError(t, err)
	assert.True(t, isExists)

	dropDbSql := fmt.Sprintf("drop database if exists %s", GetDatabaseName(url))
	_, err = pool.Query(ctx, dropDbSql)
	assert.NoError(t, err)
}
