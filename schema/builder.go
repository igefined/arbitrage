package schema

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/igdotog/core/config"
	"github.com/igdotog/core/logger"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
)

type QBuilder struct {
	pool *pgxpool.Pool
}

func New(log *logger.Logger, cfg *config.BaseConfig, lc fx.Lifecycle) *QBuilder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	CreateDatabase(ctx, log, cfg.DBUrl)

	psqlCfg, err := pgx.ParseConfigWithOptions(cfg.DBUrl, pgx.ParseConfigOptions{})
	if err != nil {
		log.Errorf("failed to parsepostgres config: %s", err)
	}

	conn, err := pgxpool.New(ctx, psqlCfg.ConnString())
	if err != nil {
		log.Errorf("failed to make postgres connection for auto create: %s", err)
	}

	if lc != nil {
		lc.Append(fx.Hook{
			OnStop: func(ctx context.Context) error {
				conn.Close()

				return nil
			},
			OnStart: func(ctx context.Context) error {
				if err = conn.Ping(ctx); err != nil {
					return fmt.Errorf("failed to ping database: %v", err)
				}

				return nil
			},
		})
	}

	return &QBuilder{conn}
}

func CreateDatabase(ctx context.Context, log *logger.Logger, url string) {
	dbName := GetDatabaseName(url)
	conn, err := pgxpool.New(ctx, ReplaceDbName(url, "postgres"))
	if err != nil {
		log.Errorf("failed to make postgres connection for auto create: %s", err)
	}
	defer conn.Close()

	var exists bool

	checkingSql := `select exists(select datname from pg_catalog.pg_database where datname = $1) as exist`
	row := conn.QueryRow(ctx, checkingSql, dbName)
	if err := row.Scan(&exists); err != nil {
		log.Errorf("autocreate db: failed to check the existence of the database: %s", err)
	}

	if exists {
		log.Infof("autocreate db: the database \"%s\" already exists", dbName)
	} else {
		if _, err := conn.Exec(ctx, fmt.Sprintf(`create database "%s"`, dbName)); err != nil {
			log.Errorf("failed to create database: %s", err)
		} else {
			log.Info("autocreate db: database created successfully")
		}
	}
}

func GetDatabaseName(url string) string {
	re := regexp.MustCompile(`(([0-9]+\/)([a-z_]+)+)`)
	out := strings.Split(re.FindString(url), "/")

	if len(out) == 2 {
		return out[1]
	}

	return ""
}

func ReplaceDbName(dbUrl, dbName string) string {
	parsed, err := url.Parse(dbUrl)
	if err != nil {
		return dbUrl
	}

	parsed.Path = "/" + dbName

	return parsed.String()

}

func (qb QBuilder) Querier() *pgxpool.Pool {
	return qb.pool
}

func (qb QBuilder) ConnString() string {
	return qb.pool.Config().ConnString()
}
