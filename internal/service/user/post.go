package user

import (
	"base/internal/models"
	hashpassword "base/internal/utils/hash_password"
	"base/internal/utils/web"
	"context"
)

func (s *userService) Create(ctx context.Context, userInput Input) (*models.User, error) {
	if usr, _ := s.userRepo.GetByEmail(ctx, userInput.Email); usr != nil {
		return nil, web.BadRequest("Tài khoản đã tồn tại")
	}

	if err := s.validator.ValidateStruct(userInput); err != nil {
		return nil, err
	}
	if userInput.Password == "" {
		return nil, web.BadRequest("Mật khẩu không được để trống")
	}

	usr := userInput.ToModel()
	pass, err := usr.Password.GenerateHashedPassword()
	if err != nil {
		return nil, err
	}
	usr.Password = hashpassword.NewPassword(pass)
	err = s.userRepo.R_Create(ctx, usr)
	return usr, err
}
