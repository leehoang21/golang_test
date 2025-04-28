package feature

import (
	"base/internal/models"
	"base/internal/utils/web"
	"context"
)

func (s *featureService) Create(ctx context.Context, input FeatureInput) (*models.Feature, error) {
	if err := s.validator.ValidateStruct(input); err != nil {
		return nil, err
	}

	if f, _ := s.repos.GetByName(ctx, input.Name); f != nil {
		return nil, web.BadRequest("feature name đã tồn tại")
	}
	if f, _ := s.repos.GetByApi(ctx, input.Api); f != nil {
		return nil, web.BadRequest("feature api đã tồn tại")
	}

	f := input.ToModel()
	err := s.repos.R_Create(ctx, f)
	if err != nil {
		return nil, err
	}
	return f, nil
}
