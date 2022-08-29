package models

import (
	"base/internal/base/model"
	hashpassword "base/internal/utils/hash_password"
)

type User struct {
	model.BaseModel `bson:",inline"`
	Email           string                    `json:"email" bson:"email"`
	Phone           string                    `json:"phone" bson:"phone"`
	Password        hashpassword.HashPassword `json:"password" bson:"password"`
	FullName        string                    `json:"fullname" bson:"fullname"`
	Birthday        string                    `json:"birthday" bson:"birthday"`
	ZaloID          string                    `json:"zalo_id" bson:"zalo_id"`
}
