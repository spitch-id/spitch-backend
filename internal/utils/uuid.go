package utils

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

func GetUUID() (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		log.Errorf("failed to generate UUID: %v", err)
		return "", err
	}

	return id.String(), err
}
