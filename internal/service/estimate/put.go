package estimate

import (
	"base/internal/models"
	"base/internal/utils/web"
	"context"
)

func (s *estimateService) Update(ctx context.Context, id string, input Input) (*models.Estimate, error) {
	if err := s.validator.ValidateStruct(input); err != nil {
		return nil, err
	}

	var rExist *models.Estimate
	if err := s.repos.R_SelectByID(ctx, id, &rExist); err != nil {
		return nil, err
	}
	if estimate, _ := s.repos.GetByTaskID(ctx, input.TaskID); estimate != nil && id != estimate.ID {
		return nil, web.BadRequest("task id đã tồn tại")
	}
	CreatedByEmail := rExist.CreatedByEmail
	baseModel := rExist.BaseModel

	rExist = input.ToModel()
	rExist.BaseModel = baseModel
	rExist.ID = id
	rExist.CreatedByEmail = CreatedByEmail
	err := s.repos.R_Update(ctx, rExist)
	return rExist, err
}
