package user

import (
	"base/internal/models"
	"context"
)

func (s *userService) GetByID(ctx context.Context, id string) (*models.User, error) {
	return s.userRepo.GetByID(ctx, id)
}
