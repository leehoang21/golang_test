package user

import (
	"base/internal/handler/filter"
	"base/internal/models"
	"base/internal/repository"
	"context"

	"base/internal/utils/validator"
)

type Service interface {
	List(ctx context.Context, input *filter.UserListParams) ([]models.User, int64, error)
	ResetPass(ctx context.Context, id string, user UserResetPassInput) (*models.User, error)
	GetByID(ctx context.Context, id string) (*models.User, error)
	Create(ctx context.Context, userInput UserInput) (*models.User, error)
	Update(ctx context.Context, id string, input UserInput) (*models.User, error)
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
