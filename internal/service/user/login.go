package user

import (
	"base/internal/models"
	"base/internal/utils/web"
	"context"
)

type LoginInput struct {
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Platform string `json:"platform"`
}

func (s *userService) Login(ctx context.Context, input LoginInput) (u *models.User, err error) {
	if input.Email != "" {
		u, err = s.userRepo.GetByEmail(ctx, input.Email)
		if err != nil {
			return nil, err
		}
	} else if input.Phone != "" {
		u, err = s.userRepo.GetByPhone(ctx, input.Phone)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, web.ErrorOK("account not found")
	}
	if err := u.Password.ComparePassword(input.Password); err != nil {
		return nil, web.Unauthorized("Mật khẩu sai")
	}
	u.Password.Set("")
	return
}
