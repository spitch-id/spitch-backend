package dto

type UserAuthRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserAuthResponse struct {
	Email string `json:"email"`
}

type User struct {
	ID              int64  `json:"id"`
	Email           string `json:"email"`
	FullName        string `json:"full_name"`
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	EmailVerifiedAt string `json:"email_verified_at"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

type UserResponse struct {
	ID              int64  `json:"id"`
	Email           string `json:"email"`
	FullName        string `json:"full_name"`
	Phone           string `json:"phone"`
	EmailVerifiedAt string `json:"email_verified_at"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}
