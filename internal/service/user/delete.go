package user

import (
	"context"
)

func (s *userService) Delete(ctx context.Context, id string) error {
	return s.userRepo.R_DeleteByID(ctx, id)
}
