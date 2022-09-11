package user

import (
	"base/internal/models"
	hashpassword "base/internal/utils/hash_password"
	"base/internal/utils/web"
	"context"
)

func (s *userService) Create(ctx context.Context, userInput UserCreateInput) (*models.User, error) {
	if usr, _ := s.userRepo.GetByEmail(ctx, userInput.Email); usr != nil {
		return nil, web.BadRequest("Tài khoản đã tồn tại")
	}
	user := &models.User{
		Email:    userInput.Email,
		Phone:    userInput.Phone,
		FullName: userInput.Fullname,
		Birthday: userInput.Birthday,
		ZaloID:   userInput.ZaloID,
	}
	var p = hashpassword.NewPassword(userInput.Password)
	pass, err := p.GererateHashedPassword()
	if err != nil {
		return nil, err
	}
	user.Password = hashpassword.NewPassword(pass)
	err = s.userRepo.R_Create(ctx, user)
	if err != nil {
		return nil, err
	}
	user.Password.Set("")
	return user, nil
}
