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
	err = godotenv.Load(filepath.Join(root, ".env"))
	if err != nil {
		return cfg, nil
	}
	cfg.Host = os.Getenv("DB_HOST")
	cfg.Port = os.Getenv("DB_PORT")
	cfg.User = os.Getenv("DB_USER")
	cfg.Password = os.Getenv("DB_PASSWORD")
	cfg.Name = os.Getenv("DB_NAME")
	fmt.Println(cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	if err = envconfig.Process("", &cfg); err != nil {
		return
	}
	return cfg, nil
}
