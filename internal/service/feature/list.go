package feature

import (
	"base/internal/handler/filter"
	"base/internal/models"
	"context"
)

func (s featureService) List(ctx context.Context, input *filter.FeatureListParams) ([]models.Feature, int64, error) {
	var res []models.Feature
	var total, err = s.repos.R_SearchAndCount(ctx, input, &res)
	return res, total, err
}
