package dto

type CreateUserInput struct {
	Username string `json:"username" binding: "required"`
	Password string `json:"password" binding: "required"`
	Name     string `json:"name"`
}

type LoginInput struct {
	Username string `json:"username" binding: "required"`
	Password string `json:"password" binding: "required"`
}

type LoginResponse struct {
	AccessToken string `json: "accessToken"`
}