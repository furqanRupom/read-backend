package middlewares

import (
	"read-backend/config"
	sql "read-backend/sql/db"
	"read-backend/tokens"
	"context"
	"net/http"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gopkg.in/mail.v2"
)

type APIContext struct {
	Tokens tokens.TokensService
	Logger *zap.Logger
	DBQueries *sql.Queries
	SMTPDialer *mail.Dialer
	Redis *redis.Client
	MainConfig config.MainConfig
}

const apiContextKey = "apiContext"

func GetAPIContext(context context.Context) *APIContext {
	return context.Value(apiContextKey).(*APIContext)
}

func ApiContextMiddleware(apiContext *APIContext, next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), apiContextKey, apiContext)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
