package task

import (
	"base/internal/models"
	"context"
)

func (s *taskService) GetByID(ctx context.Context, id string) (*models.Task, error) {
	return s.repos.GetByID(ctx, id)
}
