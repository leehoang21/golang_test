package user

import (
	"base/internal/models"
	hashpassword "base/internal/utils/hash_password"
)

type UserInput struct {
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" `
	Password string `json:"password" `
	Fullname string `json:"fullname" validate:"required"`
	ZaloID   string `json:"zalo_id"`
	Birthday string `json:"birthday"`
}

func (u UserInput) ToModel() *models.User {
	return &models.User{
		Email:    u.Email,
		Phone:    u.Phone,
		Password: hashpassword.NewPassword(u.Password),
		FullName: u.Fullname,
		Birthday: u.Birthday,
	}
}

type UserResetPassInput struct {
	PasswordOld string `json:"password_old" validate:"required"`
	PasswordNew string `json:"password_new" validate:"required"`
}
