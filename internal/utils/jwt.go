package utils

import (
	"fmt"
	"os"
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

func CreateJWT(payload JWTClaim, fileLocation string) (string, error) {
	if fileLocation == "" {
		fileLocation = "internal/certs/token/private.pem"
	}

	claim := JWTClaim{
		ID:     payload.ID,
		UserID: payload.UserID,
		Email:  payload.Email,
		RegisteredClaims: &jwt.RegisteredClaims{
			ID:        payload.ID,
			Issuer:    payload.Issuer,
			ExpiresAt: jwt.NewNumericDate(payload.ExpiresAt),
			IssuedAt:  jwt.NewNumericDate(payload.CreatedAt),
			Subject:   payload.Email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)

	privateKey, err := os.ReadFile(fileLocation)
	if err != nil {
		return "", err
	}

	parsedKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", err
	}

	jwtToken, err := token.SignedString(parsedKey)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func VerifyToken(tokenStr string, fileLocation string) (*JWTClaim, error) {
	if fileLocation == "" {
		fileLocation = "internal/certs/token/public.pem"
	}

	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		publicKey, err := os.ReadFile(fileLocation)
		if err != nil {
			return nil, err
		}

		parsedKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
		if err != nil {
			return nil, err
		}

		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("invalid token signing method")
		}

		return parsedKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid or expired token")
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
