package task

import (
	"context"
)

func (s *taskService) Delete(ctx context.Context, id string) error {
	return s.repos.R_DeleteByID(ctx, id)
}
