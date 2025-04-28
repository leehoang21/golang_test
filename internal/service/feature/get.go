package feature

import (
	"base/internal/models"
	"context"
)

func (s *featureService) GetByID(ctx context.Context, id string) (*models.Feature, error) {
	return s.repos.GetByID(ctx, id)
}
