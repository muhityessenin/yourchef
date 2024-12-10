package routes

import (
	"github.com/gin-gonic/gin"
	"recipe_service/internal/api/handler"
)

func InitRoutes(router *gin.RouterGroup, recipeHandler *handler.RecipeHandler) {
	router.POST("/r", recipeHandler.GetRecipeList)
	router.GET("/:id", recipeHandler.GetRecipeById)
	router.GET("/random", recipeHandler.GetRandomRecipes)
}
