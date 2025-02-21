package middlewares

import (
	"read-backend/config"
	sql "read-backend/sql/db"
	"read-backend/tokens"
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/cors"
	"go.uber.org/zap"
)

func CreateMiddleware(
	mainConfig config.MainConfig,
	logger *zap.Logger,
	pool *pgxpool.Pool,
	handler http.Handler,
) http.Handler {
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   mainConfig.Cors.AllowedOrigins,
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})

	redisClient := mainConfig.Redis.CreateClient()
	err := redisClient.Exists(context.Background(), "check").Err()
	if err != nil {
		logger.Fatal("Failed to connect to redis", zap.Error(err))
	}
	apiContext := &APIContext{
		Tokens: tokens.TokensService{
			Config: mainConfig.JWT,
		},
		Logger:    logger,
		DBQueries: sql.New(pool),
		Redis: redisClient,
		MainConfig: mainConfig,
	}
	return ApiContextMiddleware(
		apiContext,
		HTTPLoggingMiddleware(
			corsMiddleware.Handler(
				InjectHttpObjects(
					AuthMiddleware(
						&mainConfig.Cookie,
						handler,
					),
				),
			),
		),
	)
}
