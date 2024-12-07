package interfaces

import "recipe_service/internal/domain/recipe"

type RecipeService interface {
	GetRecipes(ingredients []string) ([]recipe.Entity, error)
	GetRecipeById(id string) (recipe.Entity, error)
}
