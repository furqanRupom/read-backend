package main

import (
	"read-backend/config"
	resolvers "read-backend/graphql/resolvers"
	middlewares "read-backend/middlewares"
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
    "github.com/99designs/gqlgen/graphql/playground"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	godotenv.Load()
	mainConfig := config.CreateMainConfig()
	logger := mainConfig.Logging.CreateLogger()

	conn, err := pgxpool.New(context.Background(), mainConfig.DB.ToURL())
	if err != nil {
		logger.Fatal("Unable to connect to database", zap.Error(err))
	}
	defer conn.Close()

	mux := http.NewServeMux()
	mux.Handle("/graphql", resolvers.CreateHandler())
	mux.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	log.Printf("connect to http://localhost:%d/ for GraphQL playground",mainConfig.Server.Port)
	server := middlewares.CreateMiddleware(mainConfig, logger, conn, mux)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(mainConfig.Server.Port), server))
}
