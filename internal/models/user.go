package models

import (
	"base/internal/base/model"
	hashpassword "base/internal/utils/hash_password"
)

type User struct {
	model.BaseModel `bson:",inline"`
	Email           string                `json:"email" bson:"email"`
	Phone           string                `json:"phone" bson:"phone"`
	Password        hashpassword.Password `json:"password,omitempty" bson:"password"`
	FullName        string                `json:"fullname" bson:"fullname"`
	Birthday        string                `json:"birthday" bson:"birthday"`
	OrgID           string                `json:"org_id" bson:"org_id"`
	ZaloID          string                `json:"zalo_id" bson:"zalo_id"`
}
