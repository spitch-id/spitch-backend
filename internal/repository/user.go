package repository

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5"
	"github.com/spitch-id/spitch-backend/domain"
	"github.com/spitch-id/spitch-backend/internal/utils"
)

type userRepository struct {
}

// ChangePassword implements domain.UserRepository.
func (u *userRepository) ChangePassword(ctx context.Context, tx pgx.Tx, user *domain.User) (*domain.User, error) {
	query := `UPDATE users SET password = $1, updated_at = $2 WHERE id = $3`
	user.UpdatedAt = time.Now()

	_, err := tx.Exec(ctx, query, user.Password, user.UpdatedAt, user.ID)
	if err != nil {
		log.Errorf("failed to change password for user %s: %v", user.ID, err)
		return nil, err
	}

	log.Infof("password changed successfully for user %s", user.ID)
	return user, nil
}

// Create implements domain.UserRepository.
func (u *userRepository) Create(ctx context.Context, tx pgx.Tx, user *domain.User) (*domain.User, error) {
	query := `INSERT INTO users (id, phone_number, email, password, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5, NOW())`
	user.CreatedAt = time.Now()

	uuid, err := utils.GetUUID()
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(ctx, query,
		uuid,
		user.Phone,
		user.Email,
		user.Password,
		user.CreatedAt,
	)
	if err != nil {
		log.Errorf("failed to create user: %v", err)
	}

	return &domain.User{
		ID:       uuid,
		FullName: user.FullName,
		Email:    user.Email,
	}, nil
}

// Delete implements domain.UserRepository.
func (u *userRepository) Delete(ctx context.Context, tx pgx.Tx, user *domain.User) (*domain.User, error) {
	panic("unimplemented")
}

// FindByEmail implements domain.UserRepository.
func (u *userRepository) FindByEmail(ctx context.Context, tx pgx.Tx, email string) (*domain.User, error) {
	var user domain.User
	query := `SELECT id, email, password FROM users WHERE email = $1`

	err := tx.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return &domain.User{}, nil
		}
		return nil, err
	}
	return &user, nil
}

// FindByID implements domain.UserRepository.
func (u *userRepository) FindByID(ctx context.Context, tx pgx.Tx, id string) (*domain.User, error) {
	query := `SELECT * FROM users WHERE id = $1`
	var user domain.User
	err := tx.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.FullName,
		&user.Phone,
		&user.Password,
		&user.EmailVerifiedAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return &domain.User{}, nil
		}
		return nil, err
	}
	return &user, nil
}

// ResetPassword implements domain.UserRepository.
func (u *userRepository) ResetPassword(ctx context.Context, tx pgx.Tx, user *domain.User) (*domain.User, error) {
	panic("unimplemented")
}

// Update implements domain.UserRepository.
func (u *userRepository) Update(ctx context.Context, tx pgx.Tx, user *domain.User) (*domain.User, error) {
	now := time.Now()
	query := `UPDATE users SET full_name = $1, phone = $2, email = $3, password = $4, updated_at = $5 WHERE id = $6`
	table, err := tx.Exec(ctx, query,
		user.FullName,
		user.Phone,
		user.Email,
		user.Password,
		now,
		user.ID,
	)
	if err != nil {
		log.Errorf("failed to update user: %v", err)
		return nil, err
	}

	if table.RowsAffected() == 0 {
		log.Warnf("no rows updated for user ID: %s", user.ID)
		return nil, fiber.NewError(fiber.StatusNotFound, "User not found")
	}
	log.Infof("user updated successfully: %s", user.ID)

	return &domain.User{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
	}, nil
}

// VerifyEmail implements domain.UserRepository.
func (u *userRepository) VerifyEmail(ctx context.Context, tx pgx.Tx, email string) (*domain.User, error) {
	panic("unimplemented")
}

func NewUserRepository() domain.UserRepository {
	return &userRepository{}
}
