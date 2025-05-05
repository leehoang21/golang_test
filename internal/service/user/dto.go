package user

import (
	"base/internal/models"
	hashpassword "base/internal/utils/hash_password"
)

type Input struct {
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" `
	Password string `json:"password" `
	Fullname string `json:"fullname" validate:"required"`
	Birthday string `json:"birthday"`
}

func (u Input) ToModel() *models.User {
	return &models.User{
		Password: hashpassword.Password(u.Password),
		Email:    u.Email,
		Phone:    u.Phone,
		FullName: u.Fullname,
		Birthday: u.Birthday,
	}
}

type ResetPassInput struct {
	PasswordOld string `json:"password_old" validate:"required"`
	PasswordNew string `json:"password_new" validate:"required"`
}
