package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_config"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_logger"
	"github.com/go-chi/chi/v5"
	chiWare "github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"plutus/internal/common/auth"
	"plutus/internal/common/handler"
	"plutus/internal/common/middleware"
	"plutus/internal/common/transaction"
	"plutus/internal/gen/db"
	"plutus/internal/user"
	"syscall"
	"time"
)

func main() {
	ctx := context.Background()

	config := hz_config.NewConfigWithAutomaticDetection()
	if err := config.Read(); err != nil {
		slog.Error("Could not read config.", "err", err)
		os.Exit(1)
	}

	logger := hz_logger.NewZeroLogAdapater(hz_logger.LogLevelFromString(config.ValueOrPanic("log.level")))

	logger.Info(ctx, "Initializing plutus api.")

	pool, err := initPostgresPool(ctx, config)
	if err != nil {
		logger.ErrorArgs(ctx, "Failed to init postgres pool: %s", err)
		os.Exit(1)
	}

	authorizer := auth.NewAuthorizer(
		logger,
		config.BoolValueOrPanic("auth.enforced"),
		config.ValueOrPanic("auth.jwt.secret"),
	)

	txMgr := transaction.NewManager(pool)
	queries := db.New(pool)

	userRepo := user.NewRdsRepository(queries, txMgr)
	userService := user.NewService(logger, txMgr, userRepo)
	userHandler := user.NewHandler(logger, authorizer, userService)

	router := initRouter(logger, config)
	handlers := []handler.PlutusHandler{
		userHandler,
	}
	for _, h := range handlers {
		hCtx := handler.Context{
			Timeout: 5 * time.Second,
			Version: handler.ApiVersionV1,
		}

		h.RegisterRoutes(router, hCtx)
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	logger.Info(ctx, "Listen and server on :8080")
	if err = server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.ErrorArgs(ctx, "Failed to listen and serve on port 8080: %w", err)
		os.Exit(1)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logger.Info(ctx, "Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.ErrorArgs(ctx, "Server forced to shutdown: %w", err)
	}

	logger.Error(ctx, "Goodbye!")
}

func initRouter(logger hz_logger.Logger, config *hz_config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.AllowCors)
	router.Use(chiWare.Recoverer)
	router.Use(chiWare.RequestID)
	router.Use(hz_logger.NewMiddleware(logger).Serve)
	router.Use(middleware.NewAuthParser(logger, config.ValueOrPanic("auth.jwt.secret")).Parse)
	return router
}

func initPostgresPool(ctx context.Context, appConfig *hz_config.Config) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(appConfig.ValueOrPanic("database.url"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse pgxpool config: %w", err)
	}

	config.MaxConns = int32(appConfig.IntValueOrPanic("database.max_connections"))
	config.MinConns = int32(appConfig.IntValueOrPanic("database.min_connections"))
	config.MaxConnLifetime = time.Duration(appConfig.IntValueOrPanic("database.max_connection_lifetime_minutes")) * time.Minute
	config.MaxConnIdleTime = time.Duration(appConfig.IntValueOrPanic("database.max_connection_idle_time_minutes")) * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to construct database pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database pool: %w", err)
	}

	return pool, nil
}
