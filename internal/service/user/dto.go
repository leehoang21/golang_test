package user

type UserCreateInput struct {
	Email    string `json:"email" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password"`
	Fullname string `json:"fullname" validate:"required"`
	ZaloID   string `json:"zalo_id"`
	Birthday string `json:"birthday"`
}

type UserResetPassInput struct {
	PasswordOld string `json:"password_old" validate:"required"`
	PasswordNew string `json:"password_new" validate:"required"`
}
