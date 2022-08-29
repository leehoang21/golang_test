package user

import (
	"base/internal/models"
	"base/internal/utils/web"
	"context"
)

func (s *userService) Create(ctx context.Context, user *models.User) (*models.User, error) {
	if usr, _ := s.userRepo.GetByEmail(ctx, user.Email); usr != nil {
		return nil, web.BadRequest("Tài khoản đã tồn tại")
	}
	pass, err := user.Password.GererateHashedPassword()
	if err != nil {
		return nil, err
	}
	user.Password.Set(pass)
	err = s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	user.Password.Set("")
	return user, nil
}
