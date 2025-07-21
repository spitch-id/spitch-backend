package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spitch-id/spitch-backend/domain"
	"github.com/spitch-id/spitch-backend/internal/dto"
	"github.com/spitch-id/spitch-backend/internal/utils"
)

type userHandler struct {
	Validate    *validator.Validate
	UserUsecase domain.UserUsecase
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
	request := new(dto.UserAuthRequest)
	if err := ctx.BodyParser(request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if err := u.Validate.Struct(request); err != nil {
		return utils.HandleValidationError(ctx, u.Validate, err)
	}

	userData, err := u.UserUsecase.GetUserByEmail(ctx.Context(), request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get user by email")
	}

	if userData == nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	if !utils.CheckPasswordHash(request.Password, userData.Password) {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid password")
	}

	response := &dto.UserAuthResponse{
		Email: userData.Email,
	}

	ctx.Status(fiber.StatusOK)
	return ctx.JSON(response)
}

// RegisterUser implements domain.UserHandler.
func (u *userHandler) RegisterUser(ctx *fiber.Ctx) error {
	request := new(dto.UserAuthRequest)
	language := ctx.Get("Accept-Language", "en")
	if language == "" {
		language = "en"
	}

	if err := ctx.BodyParser(request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if err := u.Validate.Struct(request); err != nil {
		return utils.HandleValidationError(ctx, u.Validate, err)
	}

	newUser := &dto.UserAuthRequest{
		Email:    request.Email,
		Password: request.Password,
	}

	hashedPassword, err := utils.HashPassword(newUser.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to hash password")
	}

	newUser.Password = hashedPassword

	existingUser, err := u.UserUsecase.GetUserByEmail(ctx.Context(), request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to check existing user")
	}

	if existingUser != nil && existingUser.Email == newUser.Email {
		return fiber.NewError(fiber.StatusConflict, "User already exists")
	}

	userData, err := u.UserUsecase.Register(ctx.Context(), newUser)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to register user")
	}

	if userData == nil {
		return fiber.NewError(fiber.StatusInternalServerError, "User registration failed")
	}

	response := dto.UserAuthResponse{
		Email: userData.Email,
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

func NewUserHandler(validate *validator.Validate, userUseCase domain.UserUsecase) domain.UserHandler {
	return &userHandler{
		Validate:    validate,
		UserUsecase: userUseCase,
	}
}
