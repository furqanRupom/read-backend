
package config

import (
	"errors"
	"fmt"
	"net/http"
)

type CookieConfig struct {
	SameSite http.SameSite
	Domain string
	Secure bool
}

func parseSameSite(value string) (*http.SameSite, error) {
	var val http.SameSite
	switch (value) {
	case "strict": {
		val = http.SameSiteStrictMode
		break;
	}
	case "lax": {
		val = http.SameSiteLaxMode
		break;
	}
	case "none": {
		val = http.SameSiteNoneMode
		break;
	}
	default: {
		return nil, errors.New(fmt.Sprintf("Invalid sameSite: %s", value))
	}
	}
	return &val, nil
}


func createCookieConfig() CookieConfig {
	return CookieConfig{
		SameSite: getEnv("COOKIE_SAME_SITE", false, parseSameSite),
		Secure: getEnv("COOKIE_SECURE", false, parseBool),
		Domain: getEnvWithoutParser("COOKIE_DOMAIN", true),
	}
}
