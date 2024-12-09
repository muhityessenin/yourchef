package config

import (
	"os"
)

type Config struct {
	JWTSecret        string
	UserServiceURL   string
	RecipeServiceURL string
}

func LoadConfig() *Config {
	return &Config{
		JWTSecret:        os.Getenv("JWT_SECRET"),
		UserServiceURL:   os.Getenv("USER_SERVICE_URL"),
		RecipeServiceURL: os.Getenv("RECIPE_SERVICE_URL"),
	}
}
