package config
import (
	"slices"
	"strings"
)
type corsConfig struct {
 AllowedOrigins []string
}

func parseAllowedOrigins (value string) (*[]string,error){
	val:= slices.DeleteFunc(strings.Split(value,";"), func(s string) bool { return s == "" })
	return &val, nil
}
func createCorsConfig() corsConfig {
	return corsConfig{
		AllowedOrigins: getEnv("CORS_ALLOWED_ORIGINS",false,parseAllowedOrigins),
	}
}