package recipe

type Instruction struct {
	Name  string `json:"name"`
	Steps []Step `json:"steps"`
}

type Step struct {
	Number      int          `json:"number"`
	Step        string       `json:"step"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Amount float64 `json:"amount,omitempty"`
	Unit   string  `json:"unit,omitempty"`
}
type Entity struct {
	ID                  int          `json:"id"`
	Title               string       `json:"title"`
	Description         string       `json:"description"`
	Category            string       `json:"category"`
	Matching            string       `json:"matching"`
	Ingredients         []Ingredient `json:"ingredients"`
	Instructions        []Step       `json:"instructions"`
	UsedIngredientCount int          `json:"usedIngredientCount"`
}
