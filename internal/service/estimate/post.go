package estimate

import (
	"base/internal/models"
	"base/internal/utils/web"
	"context"
)

func (s *estimateService) Create(ctx context.Context, input Input, createBy string) (*models.Estimate, error) {
	if t, _ := s.repos.GetByTaskID(ctx, input.TaskID); t != nil {
		return nil, web.BadRequest("Task id đã tồn tại")
	}

	if err := s.validator.ValidateStruct(input); err != nil {
		return nil, err
	}

	model := input.ToModel()
	model.CreatedByEmail = createBy

	err := s.repos.R_Create(ctx, model)
	return model, err
}
