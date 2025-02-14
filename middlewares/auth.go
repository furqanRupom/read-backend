package middlewares

import (
	"read-backend/config"
	"context"
	"net/http"
)

var accessTokenCookieName string = "AccessToken"
var refreshTokenCookieName string = "RefreshToken"

func getCookieValueFromRequest(request *http.Request, cookieName string) string {
	cookie, err := request.Cookie(cookieName)
	if err != nil {
		return ""
	}
	return cookie.Value
}

const cookieConfigKey = "cookieConfig"

func getCookieConfigFromContext(context context.Context) *config.CookieConfig {
	return context.Value(cookieConfigKey).(*config.CookieConfig)
}

func createAuthCookie(name string, value string, maxAge int, config *config.CookieConfig) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		MaxAge:   maxAge,
		SameSite: config.SameSite,
		Domain: config.Domain,
		HttpOnly: true,
		Secure:   config.Secure,
	}
}

func SetCookies(
	context context.Context,
	accessToken string,
	refreshToken string,
) {
	apiContext := GetAPIContext(context)
	responseWriter := getResponseWriterFromContext(context)
	http.SetCookie(
		*responseWriter,
		createAuthCookie(
			accessTokenCookieName,
			accessToken,
			int(apiContext.Tokens.Config.AccessTokenTTL.Seconds()),
			getCookieConfigFromContext(context),
		),
	)
	http.SetCookie(
		*responseWriter,
		createAuthCookie(
			refreshTokenCookieName,
			refreshToken,
			int(apiContext.Tokens.Config.RefreshTokenTTL.Seconds()),
			getCookieConfigFromContext(context),
		),
	)
}

func getResponseWriterFromContext(context context.Context) any {
	panic("unimplemented")
}

func AuthMiddleware(cookieConfig *config.CookieConfig, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := getUserFromToken(r)
		ctx := context.WithValue(
			context.WithValue(r.Context(), userKey, user),
			cookieConfigKey,
			cookieConfig,
		)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
