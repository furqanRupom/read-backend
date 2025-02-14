package config
import "fmt"

type DBConfig  struct {
  Host string 
  Port int
  User string
  Password string
  Database string

}


func (c *DBConfig) toURL() string {
  return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
    c.User,
    c.Password,
    c.Host,
    c.Port,
    c.Database,
    )
  
}


func createDBConfig () DBConfig {
  return DBConfig{
    Host: getEnvWithoutParser("POSTGRES_HOST",false),
    Port: getEnv("POSTGRES_PORT",false,parseInt),
    User: getEnvWithoutParser("POSTGRES_USER",false),
    Database: getEnvWithoutParser("POSTGRES_DB",false),
    Password: getEnvWithoutParser("POSGRES_PASSWORD",false),
  }
} 























