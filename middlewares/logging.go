package middlewares

import (
	"context"
	"net/http"
	"net/http/httputil"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

var requestIdKey = "requestId"

func GetRequestIdFromContext(ctx context.Context) string {
	return ctx.Value(requestIdKey).(string)
}

func HTTPLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiContext := GetAPIContext(r.Context())
		requestId, _ := uuid.NewUUID()
		ctx := context.WithValue(r.Context(), requestIdKey, requestId.String())
		r = r.WithContext(ctx)
		requestBytes, _ := httputil.DumpRequest(r, true)
		apiContext.Logger.Debug(
			"HTTP Request",
			zap.String("RequestId", requestId.String()),
			zap.String("Method", r.Method),
			zap.String("Path", r.URL.Path),
			zap.String("Cookie", r.Header.Get("Cookie")),
			zap.Binary("Base64", requestBytes),
		)
		lrw := NewCustomResponseWriter(w)
		next.ServeHTTP(lrw, r)
		response := lrw.GetResponse()
		responseBytes, _ := httputil.DumpResponse(&response, true)
		apiContext.Logger.Debug(
			"HTTP Response",
			zap.String("RequestId", requestId.String()),
			zap.Int("Status", lrw.statusCode),
			zap.Binary("Base64", responseBytes),
		)
	})
}
