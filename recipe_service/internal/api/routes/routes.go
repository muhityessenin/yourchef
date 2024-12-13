package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "recipe_service/docs"
	"recipe_service/internal/api/handler"
)

func InitRoutes(router *gin.RouterGroup, recipeHandler *handler.RecipeHandler) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/", recipeHandler.GetRecipeList)
	router.GET("/:id", recipeHandler.GetRecipeById)
	router.GET("/random", recipeHandler.GetRandomRecipes)
}
