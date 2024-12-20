package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	JWTSecret        string
	UserServiceURL   string
	RecipeServiceURL string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("No .env file found or can't load it: %v", err)
	}
	return &Config{
		JWTSecret:        os.Getenv("JWT_SECRET"),
		UserServiceURL:   "http://user_service:8001",
		RecipeServiceURL: "http://recipe_service:8000",
	}
}
