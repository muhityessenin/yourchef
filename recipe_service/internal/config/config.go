package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
	"path/filepath"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func LoadConfig() (cfg Config, err error) {

	root, err := os.Getwd()
	if err != nil {
		return
	}
	fmt.Println(root)
	err = godotenv.Load(filepath.Join(root, ".env"))
	if err != nil {
		return cfg, nil
	}
	cfg.Host = os.Getenv("DB_HOST")
	cfg.Port = os.Getenv("DB_PORT")
	cfg.User = os.Getenv("DB_USER")
	cfg.Password = os.Getenv("DB_PASSWORD")
	cfg.Name = os.Getenv("DB_NAME")
	if err = envconfig.Process("", &cfg); err != nil {
		return
	}
	return cfg, nil
}
