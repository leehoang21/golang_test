package estimate

import (
	"base/internal/models"
	"context"
)

func (s *estimateService) GetByID(ctx context.Context, id string) (*models.Estimate, error) {
	return s.repos.GetByID(ctx, id)
}
