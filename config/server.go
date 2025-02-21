
package config

type serverConfig struct {
	Port int
}

func createServerConfig() serverConfig {
	return serverConfig{
		Port: getEnv("PORT", false, parseInt),
	}
}
