package group_role

import (
	"base/internal/models"
	"base/internal/utils/web"
	"context"
)

func (s *groupRoleService) Create(ctx context.Context, roleInput GroupRoleInput) (*models.GroupRole, error) {
	if err := s.validator.ValidateStruct(roleInput); err != nil {
		return nil, err
	}

	if role, _ := s.repos.GetByName(ctx, roleInput.Name); role != nil {
		return nil, web.BadRequest("group role name đã tồn tại")
	}
	role := roleInput.ToModel()
	err := s.repos.R_Create(ctx, role)
	if err != nil {
		return nil, err
	}
	return role, nil
}
