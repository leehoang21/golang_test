package user

import (
	"base/internal/models"
	hashpassword "base/internal/utils/hash_password"
	"base/internal/utils/web"
	"context"
)

func (s *userService) Create(ctx context.Context, userInput models.User) (*models.User, error) {
	if usr, _ := s.userRepo.GetByEmail(ctx, userInput.Email); usr != nil {
		return nil, web.BadRequest("Tài khoản đã tồn tại")
	}

	pass, err := userInput.Password.GererateHashedPassword()
	if err != nil {
		return nil, err
	}
	userInput.Password = hashpassword.NewPassword(pass)
	err = s.userRepo.R_Create(ctx, &userInput)
	if err != nil {
		return nil, err
	}
	userInput.Password.Set("")
	return &userInput, nil
}
