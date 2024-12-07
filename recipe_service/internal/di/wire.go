//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "recipe_service/internal/api"
	"recipe_service/internal/api/handler"
	"recipe_service/internal/config"
	"recipe_service/internal/db"
	"recipe_service/internal/repository"
	"recipe_service/internal/service"
)

func Initialize(cfg config.Config) (*http.Server, error) {
	wire.Build(
		db.ConnectDatabase,
		handler.NewRecipeHandler,
		repository.NewRecipeRepository,
		service.NewRecipeService,
		http.NewServer,
	)
	return &http.Server{}, nil
}
