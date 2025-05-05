package estimate

import (
	"base/internal/models"
	"base/internal/repository"
	"context"

	"base/internal/utils/validator"
)

type Service interface {
	GetByID(ctx context.Context, id string) (*models.Estimate, error)
	Create(ctx context.Context, input Input, createBy string) (*models.Estimate, error)
	Update(ctx context.Context, id string, input Input) (*models.Estimate, error)
	Delete(ctx context.Context, id string) error
}

type estimateService struct {
	repos     repository.Estimate
	validator validator.Validator
}

func NewestimateService(repos repository.Estimate,
	validator validator.Validator) Service {
	return &estimateService{
		repos:     repos,
		validator: validator,
	}
}
