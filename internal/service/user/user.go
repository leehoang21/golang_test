package user

import (
	"context"
	"eclectric/internal/models"
	"eclectric/internal/repository"

	"eclectric/internal/utils/validator"
)

type Service interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	Update(ctx context.Context, id string, user *models.User) (*models.User, error)
	Delete(ctx context.Context, id string) error
	Login(ctx context.Context, input LoginInput) (u *models.User, err error)
}

type userService struct {
	userRepo  repository.User
	validator validator.Validator
}

func NewUserService(userRepo repository.User,
	validator validator.Validator) Service {
	return &userService{
		userRepo:  userRepo,
		validator: validator,
	}
}
