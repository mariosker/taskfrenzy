package config

import (
	"os"

	"github.com/joho/godotenv"
)

var Envs = initConfig()

type Config struct {
	PublicHost   string
	Port         string
	DBConnString string
}

func initConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost:   getEnv("PublicHost", "http://localhost"),
		Port:         getEnv("Port", "8080"),
		DBConnString: getEnv("DBConnString", "postgres://postgres:postgres@localhost/taskfrenzy"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
