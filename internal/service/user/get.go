package user

import (
	"base/internal/models"
	"context"
)

func (s *userService) GetByID(ctx context.Context, id string) (*models.User, error) {
	var u *models.User
	return u, s.userRepo.R_SelectByID(ctx, id, &u)
}
