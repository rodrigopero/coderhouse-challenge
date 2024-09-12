package dtos

type AuthorizationDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthorizationResponse struct {
	Token string `json:"token"`
}
