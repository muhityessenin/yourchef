package handler

import (
	_ "database/sql"
	_ "errors"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"net/http"
	_ "net/http"
	"recipe_service/internal/service/interfaces"
	"strconv"
)

type RecipeHandler struct {
	RecipeService interfaces.RecipeService
}

func NewRecipeHandler(service interfaces.RecipeService) *RecipeHandler {
	return &RecipeHandler{
		RecipeService: service,
	}
}
func (handler *RecipeHandler) GetRecipeList(ctx *gin.Context) {
	var input struct {
		Input []string `json:"input"`
	}

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(input.Input) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Input array is empty"})
		return
	}

	res, err := handler.RecipeService.GetRecipes(input.Input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (handler *RecipeHandler) GetRecipeById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := handler.RecipeService.GetRecipeById(strconv.FormatUint(id, 10))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (handler *RecipeHandler) GetRandomRecipes(ctx *gin.Context) {

}
