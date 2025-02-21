package middlewares

import (
	"context"
	"net/http"
)

var requestKey string = "request"
var responseWriterKey string = "ResponseWriter"

func getResponseWriterFromContext(ctx context.Context) *http.ResponseWriter {
	return ctx.Value(responseWriterKey).(*http.ResponseWriter)
}

func InjectHttpObjects(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(
			context.WithValue(r.Context(), requestKey, r),
			responseWriterKey,
			&w,
		)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
