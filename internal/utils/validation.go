package utils

import (
	"github.com/go-playground/validator/v10"
	en "github.com/go-playground/validator/v10/translations/en"
	id "github.com/go-playground/validator/v10/translations/id"
	"github.com/gofiber/fiber/v2"
	"github.com/spitch-id/spitch-backend/internal/config"
)

type ValidationErrorResponse struct {
	Message    string   `json:"message"`
	StatusCode int      `json:"statusCode"`
	Errors     []string `json:"errors"`
}

func HandleValidationError(ctx *fiber.Ctx, validate *validator.Validate, err error) error {
	if _, ok := err.(*validator.InvalidValidationError); ok {
		return ctx.Status(fiber.StatusInternalServerError).JSON(ValidationErrorResponse{
			Message:    "Invalid validation error",
			StatusCode: fiber.StatusInternalServerError,
			Errors:     []string{err.Error()},
		})
	}

	language := ctx.Get("Accept-Language", "en")
	trans, found := config.Translator.GetTranslator(language)
	if !found {
		trans, _ = config.Translator.GetTranslator("en")
	}

	switch language {
	case "id":
		_ = id.RegisterDefaultTranslations(validate, trans)
	default:
		_ = en.RegisterDefaultTranslations(validate, trans)
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(ValidationErrorResponse{
			Message:    "Validation failed",
			StatusCode: fiber.StatusBadRequest,
			Errors:     []string{err.Error()},
		})
	}

	var translated []string
	for _, ve := range validationErrors {
		translated = append(translated, ve.Translate(trans))
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(ValidationErrorResponse{
		Message:    "Validation failed",
		StatusCode: fiber.StatusBadRequest,
		Errors:     translated,
	})
}
