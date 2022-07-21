package user

import (
	"base/internal/handler/filter"
	"base/internal/model"
	"base/internal/repository"
	"context"

	"base/internal/utils/validator"
)

type UserService interface {
	List(ctx context.Context, f *filter.UserListParams) ([]model.User, error)
}

type userService struct {
	userRepo  repository.User
	validator validator.Validator
}

func NewUserService(userRepo repository.User,
	validator validator.Validator) UserService {
	return userService{
		userRepo:  userRepo,
		validator: validator,
	}
}
