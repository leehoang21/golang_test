package feature

import (
	"base/internal/models"
	"base/internal/utils/web"
	"context"
)

func (s *featureService) Update(ctx context.Context, id string, input FeatureInput) (*models.Feature, error) {
	if err := s.validator.ValidateStruct(input); err != nil {
		return nil, err
	}

	var rExist *models.Feature
	if err := s.repos.R_SelectByID(ctx, id, &rExist); err != nil {
		return nil, err
	}
	if f, _ := s.repos.GetByName(ctx, input.Name); f != nil && id != f.ID {
		return nil, web.BadRequest("feature name đã tồn tại")
	}

	rExist = input.ToModel()
	rExist.ID = id
	err := s.repos.R_Update(ctx, rExist)
	if err != nil {
		return nil, err
	}
	return rExist, nil
}
