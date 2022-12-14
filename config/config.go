package config

import (
	"fmt"
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment            string
	PostgresHost           string
	PostgresPort           int
	PostgresDatabase       string
	PostgresUser           string
	PostgresPassword       string
	LogLevel               string
	RPCPort                string
	ReviewServiceHost      string
	ReviewServicePort      int
	PostgresMigrationsPath string
}

func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "students"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "dunyo"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "1"))
	c.PostgresMigrationsPath = cast.ToString(getOrReturnDefault("POSTGRES_MIGRATIONS_PATH", "migrations"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))

	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":50051"))
	fmt.Println("config:", c)
	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
