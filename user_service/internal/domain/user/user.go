package user

type InputResponse struct {
	Name             string `json:"name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	Address          string `json:"address"`
	RegistrationDate string `json:"registration_date"`
	Role             string `json:"role"`
}

type Entity struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	Address          string `json:"address"`
	RegistrationDate string `json:"registration_date"`
	Role             string `json:"role"`
}
