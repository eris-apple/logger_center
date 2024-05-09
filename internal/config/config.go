package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	EnvKeyBindAddress = "BIND_ADDRESS"
	EnvKeyDatabaseUrl = "DATABASE_URL"
	EnvKeyRedisUrl    = "REDIS_URL"
	EnvKeyRedisPwd    = "REDIS_PASSWORD"
	EnvKeyJWTSecret   = "JWT_SECRET"
	EnvKeyHashCost    = "HASH_COST"
	EnvKeyHashSecret  = "HASH_SECRET"
)

type Config struct {
	BindAddress string `env:"BIND_ADDRESS"`
	DatabaseURL string `env:"DATABASE_URL"`
	RedisURL    string `env:"REDIS_URL"`
	RedisPwd    string `env:"REDIS_PASSWORD"`
	JWTSecret   string `env:"JWT_SECRET"`
	HashCost    int    `env:"HASH_COST"`
	HashSecret  string `env:"HASH_SECRET"`
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		EnvBindAddress = os.Getenv(EnvKeyBindAddress)
		EnvDatabaseUrl = os.Getenv(EnvKeyDatabaseUrl)
		EnvRedisUrl    = os.Getenv(EnvKeyRedisUrl)
		EnvRedisPwd    = os.Getenv(EnvKeyRedisPwd)
		EnvJWTSecret   = os.Getenv(EnvKeyJWTSecret)
		EnvHashCost    = os.Getenv(EnvKeyHashCost)
		EnvHashSecret  = os.Getenv(EnvKeyHashSecret)
	)

	if EnvBindAddress == "" {
		EnvBindAddress = "localhost:5001"
	}
	if EnvDatabaseUrl == "" {
		EnvDatabaseUrl = "postgres://root:password@localhost:5432/logger_center"
	}
	if EnvRedisUrl == "" {
		EnvRedisUrl = "localhost:3306"
	}
	if EnvJWTSecret == "" {
		EnvJWTSecret = "jwt_secret"
	}
	if EnvHashCost == "" {
		EnvHashCost = "10"
	}
	if EnvHashSecret == "" {
		EnvHashSecret = "hash_secret"
	}

	hashCost, _ := strconv.Atoi(os.Getenv("HASH_COST"))

	return &Config{
		BindAddress: EnvBindAddress,
		DatabaseURL: EnvDatabaseUrl,
		RedisURL:    EnvRedisUrl,
		RedisPwd:    EnvRedisPwd,
		JWTSecret:   EnvJWTSecret,
		HashCost:    hashCost,
		HashSecret:  EnvHashSecret,
	}
}
