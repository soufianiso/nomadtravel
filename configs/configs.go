package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost             string
	ApiGatewayPort         string
	UserMicroservicePort   string
	MoviesMicroservicePort string
	AuthMicroservicePort   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTSecret              string
	JWTExpirationInSeconds int64
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", "http://localhost"),
		ApiGatewayPort:			getEnv("PORT", "8000"),
		UserMicroservicePort:	getEnv("USER_MICROSERVICE_PORT", "50051"),
		// AuthMicroservicePort:	getEnv("AUTH_MICROSERVICE_PORT", "50052"),
		MoviesMicroservicePort:	getEnv("MOVIE_MICROSERVICE_PORT", "50052"),
		DBUser:                 getEnv("DB_USER", "postgres"),
		DBPassword:             getEnv("B_PASSWORD", "postgres"),
		DBAddress:              fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "5432")),
		DBName:                 getEnv("DB_NAME", "postgres"),
		JWTSecret:              getEnv("JWT_SECRET", "example"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600 * 24 * 7),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
