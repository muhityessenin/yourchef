package repository

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"recipe_service/internal/domain/recipe"
	"recipe_service/internal/repository/interfaces"
	"strings"
)

type Response struct {
	Offset       int             `json:"offset"`
	Number       int             `json:"number"`
	Results      []recipe.Entity `json:"results"`
	TotalResults int             `json:"totalResults"`
}
type RecipeRepository struct {
	db *sqlx.DB
}

func parseRecipes(apiResponse []byte) ([]recipe.Entity, error) {
	var recipes []recipe.Entity
	if err := json.Unmarshal(apiResponse, &recipes); err != nil {
		return nil, err
	}
	return recipes, nil
}

func (p RecipeRepository) GetRecipes(ingredients []string) ([]recipe.Entity, error) {
	ingredientsParam := strings.Join(ingredients, ",")
	apiKey := os.Getenv("API_KEY")
	url := fmt.Sprintf(
		"https://api.spoonacular.com/recipes/findByIngredients?ingredients=%s&number=5&ranking=1&apiKey=%s",
		ingredientsParam,
		apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	recipes, err := parseRecipes(body)
	if err != nil {
		log.Fatalf("Error parsing recipes: %v", err)
	}

	for i := range recipes {
		matchingPercentage := (float64(recipes[i].UsedIngredientCount) / float64(len(ingredients))) * 100
		matchingString := fmt.Sprintf("%.2f%%", matchingPercentage)
		recipes[i].Matching = matchingString
	}
	return recipes, nil
}
func (p *RecipeRepository) GetRecipeById(id string) (recipe.Entity, error) {
	url := fmt.Sprintf(
		"https://api.spoonacular.com/recipes/%s/analyzedInstructions?apiKey=%s",
		id, os.Getenv("API_KEY"))
	resp, err := http.Get(url)
	fmt.Println("Request URL:", url)
	if err != nil {
		return recipe.Entity{}, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return recipe.Entity{}, fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(bodyBytes))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return recipe.Entity{}, fmt.Errorf("error reading response body: %v", err)
	}

	entities, err := parseEntities(body)
	if err != nil {
		return recipe.Entity{}, fmt.Errorf("error parsing entities: %v", err)
	}

	if len(entities) == 0 {
		return recipe.Entity{}, fmt.Errorf("no instructions found for recipe ID %s", id)
	}

	return entities[0], nil
}

func parseEntities(apiResponse []byte) ([]recipe.Entity, error) {
	var instructions []recipe.Instruction
	if err := json.Unmarshal(apiResponse, &instructions); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	var entities []recipe.Entity

	for _, instruction := range instructions {
		entity := recipe.Entity{
			Title:        instruction.Name,
			Ingredients:  []recipe.Ingredient{},
			Instructions: []recipe.Step{},
		}

		ingredientMap := make(map[int]recipe.Ingredient)

		for _, step := range instruction.Steps {
			entity.Instructions = append(entity.Instructions, recipe.Step{
				Number: step.Number,
				Step:   step.Step,
			})

			for _, ing := range step.Ingredients {
				if _, exists := ingredientMap[ing.ID]; !exists {
					ingredientMap[ing.ID] = ing
				}
			}
		}

		for _, ing := range ingredientMap {
			entity.Ingredients = append(entity.Ingredients, recipe.Ingredient{
				Name:   ing.Name,
				Amount: ing.Amount,
				Unit:   ing.Unit,
			})
		}

		entities = append(entities, entity)
	}

	return entities, nil
}
func NewRecipeRepository(db *sqlx.DB) interfaces.RecipeRepository {
	return &RecipeRepository{db: db}
}
