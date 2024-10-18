package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Envs = initConfig()

type Config struct {
	PublicHost             string
	Port                   string
	DBConnString           string
	JWTExpirationInSeconds int64
	JWTSecret              string
}

func initConfig() Config {
	_ = godotenv.Load()

	return Config{
		PublicHost:             getEnv("PublicHost", "http://localhost"),
		Port:                   getEnv("Port", "8080"),
		DBConnString:           getEnv("DBConnString", "postgres://postgres:postgres@localhost/taskfrenzy"),
		JWTExpirationInSeconds: getEnvInt("JWTExpirationInSeconds", 3600*24*7),
		JWTSecret:              getEnv("JWTSecret", "secret"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}

	return fallback
}
