package user

import (
	"base/internal/models"
	hashpassword "base/internal/utils/hash_password"
	"base/internal/utils/web"
	"context"
	"fmt"
)

type LoginInput struct {
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Platform string `json:"platform"`
}

func (s *userService) Login(ctx context.Context, input LoginInput) (u *models.User, err error) {
	fmt.Println(input.Email)
	if input.Email != "" {
		u, err = s.userRepo.GetByEmail(ctx, input.Email)
		if err != nil {
			return nil, web.Unauthorized("Account not found" + err.Error())
		}
	} else if input.Phone != "" {
		u, err = s.userRepo.GetByPhone(ctx, input.Phone)
		if err != nil {
			return nil, web.Unauthorized("Account not found" + err.Error())
		}
	} else {
		return nil, web.Unauthorized("Account not found")
	}
	if err := u.Password.ComparePassword(hashpassword.NewPassword(input.Password)); err != nil {
		return nil, web.Unauthorized("Mật khẩu sai")
	}
	u.Password = ""
	return
}
