package usecase

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spitch-id/spitch-backend/domain"
	"github.com/spitch-id/spitch-backend/internal/dto"
)

type userUseCase struct {
	DB             *pgxpool.Pool
	UserRepository domain.UserRepository
}

// ChangePassword implements domain.UserUsecase.
func (u *userUseCase) ChangePassword(ctx context.Context, user *dto.UserAuthRequest) (*dto.UserAuthResponse, error) {
	panic("unimplemented")
}

// DeleteUser implements domain.UserUsecase.
func (u *userUseCase) DeleteUser(ctx context.Context, user *dto.UserAuthRequest) (*dto.UserAuthResponse, error) {
	panic("unimplemented")
}

// GetUserByEmail implements domain.UserUsecase.
func (u *userUseCase) GetUserByEmail(ctx context.Context, user *dto.UserAuthRequest) (*domain.User, error) {
	tx, err := u.DB.Begin(ctx)
	if err != nil {
		log.Errorf("failed to begin transaction: %v", err)
		return nil, err
	}
	defer tx.Rollback(ctx)

	userData, err := u.UserRepository.FindByEmail(ctx, tx, user.Email)
	if err != nil {
		log.Errorf("failed to get user by email: %v", err)
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		log.Errorf("failed to commit transaction: %v", err)
		return nil, err
	}

	response := &domain.User{
		Email:    userData.Email,
		ID:       userData.ID,
		Password: userData.Password,
	}

	return response, nil
}

// GetUserByID implements domain.UserUsecase.
func (u *userUseCase) GetUserByID(ctx context.Context, user *dto.UserAuthRequest) (*domain.User, error) {
	panic("unimplemented")
}

// ListUsers implements domain.UserUsecase.
func (u *userUseCase) ListUsers(ctx context.Context, user *dto.UserAuthRequest) ([]*dto.UserAuthResponse, error) {
	panic("unimplemented")
}

// Login implements domain.UserUsecase.
func (u *userUseCase) Login(ctx context.Context, user *dto.UserAuthRequest) (*dto.UserAuthResponse, error) {
	panic("unimplemented")
}

// Register implements domain.UserUsecase.
func (u *userUseCase) Register(ctx context.Context, user *dto.UserAuthRequest) (*domain.User, error) {
	tx, err := u.DB.Begin(ctx)
	if err != nil {
		log.Errorf("failed to begin transaction: %v", err)
		return nil, err
	}
	defer tx.Rollback(ctx)

	newUser := &domain.User{
		Email:    user.Email,
		Password: user.Password,
	}

	userData, err := u.UserRepository.Create(ctx, tx, newUser)
	if err != nil {
		log.Errorf("failed to get user by email: %v", err)
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		log.Errorf("failed to commit transaction: %v", err)
		return nil, err
	}

	response := &domain.User{
		Email: userData.Email,
		ID:    userData.ID,
	}

	log.Infof("user registered successfully: %s", userData.Email)
	return response, nil
}

// ResetPassword implements domain.UserUsecase.
func (u *userUseCase) ResetPassword(ctx context.Context, user *dto.UserAuthRequest) (*dto.UserAuthResponse, error) {
	panic("unimplemented")
}

// UpdateUser implements domain.UserUsecase.
func (u *userUseCase) UpdateUser(ctx context.Context, user *dto.UserAuthRequest) (*dto.UserAuthResponse, error) {
	panic("unimplemented")
}

// VerifyEmail implements domain.UserUsecase.
func (u *userUseCase) VerifyEmail(ctx context.Context, user *dto.UserAuthRequest) (*dto.UserAuthResponse, error) {
	panic("unimplemented")
}

func NewUserUseCase(db *pgxpool.Pool, userRepo domain.UserRepository) domain.UserUsecase {
	return &userUseCase{
		DB:             db,
		UserRepository: userRepo,
	}
}
