package task

import (
	"base/internal/handler/filter"
	"base/internal/models"
	"context"
)

func (s taskService) List(ctx context.Context, input *filter.TaskListParams) ([]models.Task, int64, error) {
	var res []models.Task
	var total, err = s.repos.R_SearchAndCount(ctx, input, &res)
	return res, total, err
}
