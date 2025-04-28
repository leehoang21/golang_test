package feature

import (
	"base/internal/handler/filter"
	"base/internal/models"
	"base/internal/repository"
	"context"

	"base/internal/utils/validator"
)

type Service interface {
	List(ctx context.Context, input *filter.FeatureListParams) ([]models.Feature, int64, error)
	GetByID(ctx context.Context, id string) (*models.Feature, error)
	Create(ctx context.Context, roleInput FeatureInput) (*models.Feature, error)
	Update(ctx context.Context, id string, roleInput FeatureInput) (*models.Feature, error)
	Delete(ctx context.Context, id string) error
}

type featureService struct {
	repos     repository.Feature
	validator validator.Validator
}

func NewFeatureService(repos repository.Feature,
	validator validator.Validator) Service {
	return &featureService{
		repos:     repos,
		validator: validator,
	}
}
