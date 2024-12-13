package handler

import (
	_ "database/sql"
	_ "errors"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"net/http"
	_ "net/http"
	_ "recipe_service/internal/domain/recipe"
	"recipe_service/internal/service/interfaces"
	"recipe_service/pkg"
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

// GetRecipeList godoc
// @Summary      Get Recipes by Ingredients
// @Description  Returns a list of recipes based on the provided ingredients
// @Tags         recipes
// @Accept       json
// @Produce      json
// @Param        input  body      InputIngredient  true  "Array of ingredients, e.g., {\"ingredients\": [\"tomato\", \"cheese\"]}"
// @Success      200    {object}  Response         "Success response with a list of recipes"
// @Failure      400    {object}  Response         "Bad Request: Invalid input or empty array"
// @Failure      500    {object}  Response         "Internal Server Error"
// @Router       /recipe [post]
func (handler *RecipeHandler) GetRecipeList(ctx *gin.Context) {
	var input struct {
		Input []string `json:"input"`
	}
	if err := ctx.BindJSON(&input); err != nil {
		res := pkg.Response{Status: http.StatusBadRequest, Message: err.Error(), Data: ""}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if len(input.Input) == 0 {
		res := pkg.Response{Status: http.StatusBadRequest, Message: "input is empty", Data: ""}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res, err := handler.RecipeService.GetRecipes(input.Input)
	if err != nil {
		res := pkg.Response{Status: http.StatusBadRequest, Message: err.Error(), Data: ""}
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	response := pkg.Response{Status: http.StatusOK, Data: res}
	ctx.JSON(http.StatusOK, response)
}

// GetRecipeById godoc
// @Summary      Get Recipe By ID
// @Description  Returns a recipe with the specified ID
// @Tags         recipes
// @Produce      json
// @Param        id   path      int  true  "Recipe ID"
// @Success      200  {object}  recipe.Entity   "Details of the recipe"
// @Failure      400  {object}  Response
// @Failure      500  {object}  Response
// @Router       /recipe/{id} [get]
func (handler *RecipeHandler) GetRecipeById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response := pkg.Response{Status: http.StatusBadRequest, Message: err.Error(), Data: ""}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	res, err := handler.RecipeService.GetRecipeById(strconv.FormatUint(id, 10))
	if err != nil {
		response := pkg.Response{Status: http.StatusBadRequest, Message: err.Error(), Data: ""}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := pkg.Response{Status: http.StatusOK, Data: res}
	ctx.JSON(http.StatusOK, response)
}

// GetRandomRecipes godoc
// @Summary      Get Random Recipes
// @Description  Returns a list of random recipes
// @Tags         recipes
// @Produce      json
// @Success      200 {object} []recipe.Entity  "List of random recipes"
// @Failure      500 {object} Response
// @Router       /recipe/random [get]
func (handler *RecipeHandler) GetRandomRecipes(ctx *gin.Context) {
	res, err := handler.RecipeService.GetRandomRecipes()
	if err != nil {
		response := pkg.Response{Status: http.StatusBadRequest, Message: err.Error(), Data: ""}
		ctx.JSON(http.StatusInternalServerError, response)
	}
	response := pkg.Response{Status: http.StatusOK, Data: res}
	ctx.JSON(http.StatusOK, response)
}
