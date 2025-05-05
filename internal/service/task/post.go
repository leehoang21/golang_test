package task

import (
	"base/internal/models"
	"base/internal/utils/web"
	"context"
)

func (s *taskService) Create(ctx context.Context, input Input, createBy string) (*models.Task, error) {
	if t, _ := s.repos.GetByTitle(ctx, input.Title); t != nil {
		return nil, web.BadRequest("Task title đã tồn tại")
	}

	if err := s.validator.ValidateStruct(input); err != nil {
		return nil, err
	}

	t := input.ToModel()
	t.CreatedByEmail = createBy

	err := s.repos.R_Create(ctx, t)
	if err != nil {
		return nil, err
	}
	return t, nil
}
