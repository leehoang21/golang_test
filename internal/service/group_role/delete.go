package group_role

import (
	"context"
)

func (s *groupRoleService) Delete(ctx context.Context, id string) error {
	return s.repos.R_DeleteByID(ctx, id)
}
