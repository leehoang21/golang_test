package group_role

import (
	"base/internal/models"
	"context"
)

func (s *groupRoleService) GetByID(ctx context.Context, id string) (*models.GroupRole, error) {
	return s.repos.GetByID(ctx, id)
}
