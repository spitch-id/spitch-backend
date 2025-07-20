package dto

type UserAuthRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserAuthResponse struct {
	Email string `json:"email"`
}
