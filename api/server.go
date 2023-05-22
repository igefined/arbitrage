package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/igdotog/core/config"
	"github.com/igdotog/core/logger"
	"github.com/igilgyrg/arbitrage/api/endpoints"

	"go.uber.org/fx"
)

type Server struct {
	mux *http.ServeMux

	cfg       *config.BaseConfig
	logger    *logger.Logger
	endpoints endpoints.Endpoint
}

func NewServer(
	cfg *config.BaseConfig,
	logger *logger.Logger,
	endpoints endpoints.Endpoint,
	lc fx.Lifecycle,
) *Server {
	mux := http.NewServeMux()

	srv := &Server{
		cfg:       cfg,
		logger:    logger,
		endpoints: endpoints,

		mux: mux,
	}

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: mux,
	}

	srv.mapRoutes()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Infof("http server start at %s", httpServer.Addr)
			go func() {
				if err := httpServer.ListenAndServe(); err != nil {
					logger.Error(err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return httpServer.Shutdown(ctx)
		},
	})

	return srv
}
