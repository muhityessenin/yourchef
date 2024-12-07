package interfaces

import "recipe_service/internal/domain/recipe"

type RecipeRepository interface {
	GetRecipes(ingredients []string) ([]recipe.Entity, error)
	GetRecipeById(id string) (recipe.Entity, error)
}
