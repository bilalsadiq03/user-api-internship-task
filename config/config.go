package config

import "os"


type Config struct {
	DBurl string
	Port  string
}


func Load() *Config {
	return &Config{
		DBurl: getEnv("DBURL", "postgres://postgres:9528296572@localhost:5432/user_api?sslmode=disable"),
		Port:  getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}	
	return fallback
}