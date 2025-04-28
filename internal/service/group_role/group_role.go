package group_role

import (
	"base/internal/handler/filter"
	"base/internal/models"
	"base/internal/repository"
	"context"

	"base/internal/utils/validator"
)

type Service interface {
	List(ctx context.Context, input *filter.GroupRoleListParams) ([]models.GroupRole, int64, error)
	GetByID(ctx context.Context, id string) (*models.GroupRole, error)
	Create(ctx context.Context, roleInput GroupRoleInput) (*models.GroupRole, error)
	Update(ctx context.Context, id string, roleInput GroupRoleInput) (*models.GroupRole, error)
	Delete(ctx context.Context, id string) error
}

type groupRoleService struct {
	repos     repository.GroupRole
	validator validator.Validator
}

func NewGroupRoleService(repos repository.GroupRole,
	validator validator.Validator) Service {
	return &groupRoleService{
		repos:     repos,
		validator: validator,
	}
}
