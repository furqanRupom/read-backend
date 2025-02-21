package config

type MainConfig struct {
	DB      DBConfig
	Cors    corsConfig
	Logging LoggingConfig
	JWT     JWTConfig
	Redis   RedisConfig
	Cookie  CookieConfig
	Server  serverConfig
}

func CreateMainConfig() MainConfig {
	return MainConfig{
		Cookie:  createCookieConfig(),
		Cors:    createCorsConfig(),
		DB:      createDBConfig(),
		Server:  createServerConfig(),
		JWT:     createJWTConfig(),
		Redis:   createRedisConfig(),
		Logging: createLoggingConfig(),
	}
}
