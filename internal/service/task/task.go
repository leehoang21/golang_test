package task

import (
	"base/internal/handler/filter"
	"base/internal/models"
	"base/internal/repository"
	"context"

	"base/internal/utils/validator"
)

type Service interface {
	List(ctx context.Context, input *filter.TaskListParams) ([]models.Task, int64, error)
	GetByID(ctx context.Context, id string) (*models.Task, error)
	Create(ctx context.Context, input Input, createBy string) (*models.Task, error)
	Update(ctx context.Context, id string, input Input) (*models.Task, error)
	Delete(ctx context.Context, id string) error
}

type taskService struct {
	repos     repository.Task
	validator validator.Validator
}

func NewtaskService(repos repository.Task,
	validator validator.Validator) Service {
	return &taskService{
		repos:     repos,
		validator: validator,
	}
}
