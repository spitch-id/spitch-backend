package dto

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaim struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	Issuer    string    `json:"-"`
	ExpiresAt time.Time `json:"-"`
	CreatedAt time.Time `json:"-"`
	*jwt.RegisteredClaims
}
