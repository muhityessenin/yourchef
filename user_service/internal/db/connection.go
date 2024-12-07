package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"user_service/internal/config"
)

func ConnectDatabase(cfg config.Config) (store *sqlx.DB, err error) {
	psqlUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		"postgres", cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	log.Printf("url: %v", psqlUrl)

	store, err = sqlx.Connect("postgres", psqlUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return
	}
	log.Printf("Connected to database with URL: %s", psqlUrl)
	return

}
