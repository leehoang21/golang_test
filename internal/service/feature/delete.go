package feature

import (
	"context"
)

func (s *featureService) Delete(ctx context.Context, id string) error {
	return s.repos.R_DeleteByID(ctx, id)
}
