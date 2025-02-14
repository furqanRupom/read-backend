package middlewares

import (
	"read-backend/tokens"
	"context"
	"net/http"

)

func refresh(refreshToken string) (string, error) {
	return refreshToken, nil
}

func parseAccessToken(token string, context context.Context) (*tokens.JWTUser, error) {
	apiContext := GetAPIContext(context)
	return apiContext.Tokens.ParseAccessToken(token)
}

func getAccessTokenFromRefreshToken(request *http.Request) string {
	refreshToken := getCookieValueFromRequest(request, refreshTokenCookieName)
	if refreshToken == "" {
		return ""
	}
	accessToken, err := refresh(refreshToken)
	if err != nil {
		return ""
	}
	return accessToken
}

func getUserFromToken(request *http.Request) *tokens.JWTUser {
	accessToken := getCookieValueFromRequest(request, accessTokenCookieName)
	if accessToken == "" {
		accessToken = getAccessTokenFromRefreshToken(request)
	}
	if accessToken == "" {
		return nil
	}
	user, err := parseAccessToken(accessToken, request.Context())
	if err != nil {
		return nil
	}
	return user
}

var userKey string = "user"

func GetUserFromContext(ctx context.Context) *tokens.JWTUser {
	return ctx.Value(userKey).(*tokens.JWTUser)
}
