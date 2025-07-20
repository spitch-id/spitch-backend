package domain

import (
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID                string       `db:"id"`
	FullName          string       `db:"full_name"`
	Phone             string       `db:"phone"`
	Email             string       `db:"email"`
	Password          string       `db:"password"`
	EmailVerifiedAtDB sql.NullTime `db:"email_verified_at"`
	EmailVerifiedAt   time.Time    `db:"-"`
	CreatedAt         time.Time    `db:"created_at"`
	UpdatedAt         time.Time    `db:"updated_at"`
}

type UserRepository interface {
	FindByID(id string) (*User, error)
	FindByEmail(email string) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id string) error
	VerifyEmail(id string) error
	ChangePassword(id, newPassword string) error
	ResetPassword(id, newPassword string) error
}

type UserUsecase interface {
	Register(user *User) error
	Login(email, password string) (*User, error)
	VerifyEmail(id string) error
	ChangePassword(id, oldPassword, newPassword string) error
	ResetPassword(email, newPassword string) error
	GetUserByID(id string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id string) error
	ListUsers() ([]*User, error)
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
