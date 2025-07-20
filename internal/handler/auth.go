package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spitch-id/spitch-backend/domain"
	"github.com/spitch-id/spitch-backend/internal/dto"
)

type userHandler struct {
	Validate *validator.Validate
}

// ChangePassword implements domain.UserHandler.
func (u *userHandler) ChangePassword(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// DeleteUser implements domain.UserHandler.
func (u *userHandler) DeleteUser(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// GetUserByEmail implements domain.UserHandler.
func (u *userHandler) GetUserByEmail(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// GetUserByID implements domain.UserHandler.
func (u *userHandler) GetUserByID(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// ListUsers implements domain.UserHandler.
func (u *userHandler) ListUsers(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// LoginUser implements domain.UserHandler.
func (u *userHandler) LoginUser(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// RegisterUser implements domain.UserHandler.
func (u *userHandler) RegisterUser(ctx *fiber.Ctx) error {
	request := new(dto.UserAuthRequest)
	if err := ctx.BodyParser(request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	if err := u.Validate.Struct(request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Validation failed: "+err.Error())
	}

	response := dto.UserAuthResponse{
		Email: request.Email,
	}

	ctx.Status(fiber.StatusCreated)
	return ctx.JSON(response)
}

// ResetPassword implements domain.UserHandler.
func (u *userHandler) ResetPassword(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// UpdateUser implements domain.UserHandler.
func (u *userHandler) UpdateUser(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// VerifyEmail implements domain.UserHandler.
func (u *userHandler) VerifyEmail(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

func NewUserHandler(validate *validator.Validate) domain.UserHandler {
	return &userHandler{
		Validate: validate,
	}
}
