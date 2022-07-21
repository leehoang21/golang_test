package model

import "base/internal/base/model"

type User struct {
	model.BaseModel `bson:",inline"`
	Email           string `json:"email" bson:"email"`
	Phone           string `json:"phone" bson:"phone"`
	Password        string `json:"password" bson:"password"`
	FirstName       string `json:"first_name" bson:"first_name"`
	LastName        string `json:"last_name" bson:"last_name"`
	FullName        string `json:"fullname" bson:"fullname"`
	Birthday        string `json:"birthday" bson:"birthday"`
}
