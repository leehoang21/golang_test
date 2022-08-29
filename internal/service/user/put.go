package user

import (
	"base/internal/models"
	"base/internal/utils/web"
	"context"
)

func (s *userService) Update(ctx context.Context, id string, user *models.User) (*models.User, error) {
	var uExist *models.User
	if err := s.userRepo.SelectByID(ctx, id, &uExist); err != nil {
		return nil, err
	}
	if uEmail, _ := s.userRepo.GetByEmail(ctx, user.Email); uEmail != nil && id != uEmail.Email {
		return nil, web.BadRequest("Email đã tồn tại")
	}
	if user.Password.String() != "" {
		pass, err := user.Password.GererateHashedPassword()
		if err != nil {
			return nil, err
		}
		user.Password.Set(pass)
	}
	user.ID = id
	err := s.userRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
