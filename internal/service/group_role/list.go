package group_role

import (
	"base/internal/handler/filter"
	"base/internal/models"
	"context"
)

func (s groupRoleService) List(ctx context.Context, input *filter.GroupRoleListParams) ([]models.GroupRole, int64, error) {
	var res []models.GroupRole
	var total, err = s.repos.R_SearchAndCount(ctx, input, &res)
	return res, total, err
}
