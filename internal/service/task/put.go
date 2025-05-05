package task

import (
	"base/internal/models"
	"base/internal/utils/web"
	"context"
)

func (s *taskService) Update(ctx context.Context, id string, input Input) (*models.Task, error) {
	if err := s.validator.ValidateStruct(input); err != nil {
		return nil, err
	}

	var rExist *models.Task
	if err := s.repos.R_SelectByID(ctx, id, &rExist); err != nil {
		return nil, err
	}
	if title, _ := s.repos.GetByTitle(ctx, input.Title); title != nil && id != title.ID {
		return nil, web.BadRequest("Task title đã tồn tại")
	}
	createByEmail := rExist.CreatedByEmail
	baseModel := rExist.BaseModel

	rExist = input.ToModel()
	rExist.ID = id
	rExist.CreatedByEmail = createByEmail
	rExist.BaseModel = baseModel
	err := s.repos.R_Update(ctx, rExist)
	if err != nil {
		return nil, err
	}
	return rExist, nil
}
