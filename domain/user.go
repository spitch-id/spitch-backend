package domain

import (
	"database/sql"
	"time"
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
	RegisterUser(user *User) (*User, error)
	LoginUser(email, password string) (*User, error)
	VerifyUserEmail(id string) error
	ChangeUserPassword(id, oldPassword, newPassword string) error
	ResetUserPassword(email, newPassword string) error
	GetUserByID(id string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id string) error
	ListUsers() ([]*User, error)
}

// UserValidator defines the methods for validating user data.
type UserValidator interface {
	ValidateRegistration(user *User) error
	ValidateLogin(email, password string) error
	ValidateEmail(email string) error
	ValidatePassword(password string) error
	ValidateUpdate(user *User) error
	ValidateDeletion(id string) error
}
