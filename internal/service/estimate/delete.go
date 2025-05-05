package estimate

import (
	"context"
)

func (s *estimateService) Delete(ctx context.Context, id string) error {
	return s.repos.R_DeleteByID(ctx, id)
}
