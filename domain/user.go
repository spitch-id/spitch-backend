package domain

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/spitch-id/spitch-backend/internal/dto"
)

type User struct {
	ID              string    `db:"id"`
	FullName        string    `db:"full_name"`
	Phone           string    `db:"phone"`
	Email           string    `db:"email"`
	Password        string    `db:"password"`
	EmailVerifiedAt time.Time `db:"email_verified_at"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

type UserRepository interface {
	FindByID(ctx context.Context, tx pgx.Tx, id string) (*User, error)
	FindByEmail(ctx context.Context, tx pgx.Tx, email string) (*User, error)
	Create(ctx context.Context, tx pgx.Tx, user *User) (*User, error)
	Update(ctx context.Context, tx pgx.Tx, user *User) (*User, error)
	Delete(ctx context.Context, tx pgx.Tx, user *User) (*User, error)
	VerifyEmail(ctx context.Context, tx pgx.Tx, email string) (*User, error)
	ChangePassword(ctx context.Context, tx pgx.Tx, user *User) (*User, error)
	ResetPassword(ctx context.Context, tx pgx.Tx, user *User) (*User, error)
}

type UserUsecase interface {
	Register(ctx context.Context, user *dto.UserAuthRequest) (*User, error)
	Login(ctx context.Context, user *dto.UserAuthRequest) (*dto.UserAuthResponse, error)
	VerifyEmail(ctx context.Context, user *dto.UserAuthRequest) (*dto.UserAuthResponse, error)
	ChangePassword(ctx context.Context, user *dto.UserAuthRequest) (*dto.UserAuthResponse, error)
	ResetPassword(ctx context.Context, user *dto.UserAuthRequest) (*dto.UserAuthResponse, error)
	GetUserByID(ctx context.Context, user *dto.UserAuthRequest) (*User, error)
	GetUserByEmail(ctx context.Context, user *dto.UserAuthRequest) (*User, error)
	UpdateUser(ctx context.Context, user *dto.UserAuthRequest) (*dto.UserAuthResponse, error)
	DeleteUser(ctx context.Context, user *dto.UserAuthRequest) (*dto.UserAuthResponse, error)
	ListUsers(ctx context.Context, user *dto.UserAuthRequest) ([]*dto.UserAuthResponse, error) //([]*User, error)
}

type UserHandler interface {
	RegisterUser(ctx *fiber.Ctx) error
	LoginUser(ctx *fiber.Ctx) error
	VerifyEmail(ctx *fiber.Ctx) error
	ChangePassword(ctx *fiber.Ctx) error
	ResetPassword(ctx *fiber.Ctx) error
	GetUserByID(ctx *fiber.Ctx) error
	GetUserByEmail(ctx *fiber.Ctx) error
	UpdateUser(ctx *fiber.Ctx) error
	DeleteUser(ctx *fiber.Ctx) error
	ListUsers(ctx *fiber.Ctx) error
}

type UserValidator interface {
	ValidateRegistration(user *User) error
	ValidateLogin(email, password string) error
	ValidateEmail(email string) error
	ValidatePassword(password string) error
	ValidateUpdate(user *User) error
	ValidateDeletion(id string) error
}
