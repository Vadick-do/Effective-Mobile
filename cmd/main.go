package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	core_config "github.com/Vadick-do/Effective-Mobile/internal/core/config"
	core_logger "github.com/Vadick-do/Effective-Mobile/internal/core/logger"
	core_pgx_pool "github.com/Vadick-do/Effective-Mobile/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "github.com/Vadick-do/Effective-Mobile/internal/core/transport/http/middleware"
	core_http_server "github.com/Vadick-do/Effective-Mobile/internal/core/transport/http/server"
	subscriptions_postgres_repository "github.com/Vadick-do/Effective-Mobile/internal/features/repository/postgres"
	subscriptions_service "github.com/Vadick-do/Effective-Mobile/internal/features/service"
	subscriptions_transport_http "github.com/Vadick-do/Effective-Mobile/internal/features/transport/http"
	"go.uber.org/zap"
)

func main() {
	cfg := core_config.NewConfigMust()
	time.Local = cfg.TimeZone
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("Failed to init application logger: ", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("application time zone", zap.Any("zone", time.Local))

	logger.Debug("initializing postgres connection pool")

	pool, err := core_pgx_pool.NewPool(ctx, core_pgx_pool.NewConfigMust())
	if err != nil {
		logger.Fatal("failed to init postgres connection pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("initializing feature", zap.String("feature", "subscriptions"))
	subscriptionsRepository := subscriptions_postgres_repository.NewSubscriptionsRepository(pool)
	subscriptionsService := subscriptions_service.NewSubscriptionsService(subscriptionsRepository)
	subscriptionsTransportHTTP := subscriptions_transport_http.NewSubscriptionsHTTPHandler(subscriptionsService)

	logger.Debug("initializing HTTP Server")
	httpConfig := core_http_server.NewConfigMust()
	httpServer := core_http_server.NewHTTPServer(
		httpConfig,
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes(subscriptionsTransportHTTP.Routes()...)

	// apiVersionRouterv2 := core_http_server.NewAPIVersionRouter(
	// 	core_http_server.ApiVersion2,
	// 	core_http_middleware.Dummy("api v2 middleware"),
	// )
	// apiVersionRouterv2.RegisterRoutes(usersTransportHTTP.Routes()...)

	httpServer.RegisterAPIRouters(
		apiVersionRouter,
		// apiVersionRouterv2,
	)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
