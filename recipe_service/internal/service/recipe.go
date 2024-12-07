package service

import (
	"recipe_service/internal/domain/recipe"
	interf "recipe_service/internal/repository/interfaces"
	services "recipe_service/internal/service/interfaces"
)

type RecipeService struct {
	RecipeRepository interf.RecipeRepository
}

func (p RecipeService) GetRecipes(ingredients []string) ([]recipe.Entity, error) {
	return p.RecipeRepository.GetRecipes(ingredients)
}

func (p RecipeService) GetRecipeById(id string) (recipe.Entity, error) {
	return p.RecipeRepository.GetRecipeById(id)
}

func NewRecipeService(recipeRepository interf.RecipeRepository) services.RecipeService {
	return &RecipeService{RecipeRepository: recipeRepository}
}
