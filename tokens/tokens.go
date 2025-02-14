package tokens

import (
	"read-backend/config"
	"errors"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTUser struct {
	Id      uuid.UUID
	IsAdmin bool
}

type TokensService struct {
	Config config.JWTConfig
}

type JWTClaims struct {
	KeyPairId int8 `json:"keyPairId"`
	jwt.RegisteredClaims
}

func (service TokensService) CreateJWTTokens(user JWTUser) (string, string, error) {
	audience := "user"
	if user.IsAdmin {
		audience = "admin"
	}
	now := time.Now()
	accessToken, err := jwt.NewWithClaims(
		&jwt.SigningMethodEd25519{}, JWTClaims{
			KeyPairId: service.Config.PrivateKey.Id,
			RegisteredClaims: jwt.RegisteredClaims{
				Audience:  jwt.ClaimStrings{audience},
				Subject:   user.Id.String(),
				ExpiresAt: jwt.NewNumericDate(now.Add(service.Config.AccessTokenTTL)),
				IssuedAt:  jwt.NewNumericDate(now),
			},
		},
	).SignedString(service.Config.PrivateKey.Key)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := jwt.NewWithClaims(
		&jwt.SigningMethodEd25519{}, JWTClaims{
			KeyPairId: service.Config.PrivateKey.Id,
			RegisteredClaims: jwt.RegisteredClaims{
				Audience:  jwt.ClaimStrings{audience},
				Subject:   user.Id.String(),
				ExpiresAt: jwt.NewNumericDate(now.Add(service.Config.RefreshTokenTTL)),
				IssuedAt:  jwt.NewNumericDate(now),
			},
		},
	).SignedString(service.Config.PrivateKey.Key)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (service TokensService) ParseAccessToken(token string) (*JWTUser, error) {
	parsedToken, err := jwt.ParseWithClaims(
		token,
		&JWTClaims{},
		func(t *jwt.Token) (interface{}, error) {
			keyPairId := t.Claims.(*JWTClaims).KeyPairId
			publicKey, ok := service.Config.PublicKeys[keyPairId]
			if !ok {
				return nil, errors.New(
					fmt.Sprintf("No public key found for id: %d", keyPairId),
				)
			}
			return publicKey, nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims := parsedToken.Claims.(*JWTClaims)
	audience := claims.Audience[0]
	isAdmin := false
	if audience == "admin" {
		isAdmin = true
	}
	return &JWTUser{Id: uuid.MustParse(claims.Subject), IsAdmin: isAdmin}, nil
}
