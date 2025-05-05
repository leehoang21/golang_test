package group_role

import (
	"base/internal/models"
	"base/internal/utils/web"
	"context"
)

func (s *groupRoleService) Update(ctx context.Context, id string, roleInput GroupRoleInput) (*models.GroupRole, error) {
	if err := s.validator.ValidateStruct(roleInput); err != nil {
		return nil, err
	}

	var rExist *models.GroupRole
	if err := s.repos.R_SelectByID(ctx, id, &rExist); err != nil {
		return nil, err
	}
	if name, _ := s.repos.GetByName(ctx, roleInput.Name); name != nil && id != name.ID {
		return nil, web.BadRequest("group role name đã tồn tại")
	}

	baseModel := rExist.BaseModel
	rExist = roleInput.ToModel()
	rExist.BaseModel = baseModel
	rExist.ID = id
	err := s.repos.R_Update(ctx, rExist)
	if err != nil {
		return nil, err
	}
	return rExist, nil
}
