//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "user_service/internal/api"
	"user_service/internal/api/handler"
	"user_service/internal/config"
	"user_service/internal/db"
	"user_service/internal/repository"
	"user_service/internal/service"
)

func Initialize(cfg config.Config) (*http.Server, error) {
	wire.Build(
		db.ConnectDatabase,
		handler.NewUserHandler,
		repository.NewUserRepository,
		service.NewUserService,
		http.NewServer,
	)
	return &http.Server{}, nil
}
