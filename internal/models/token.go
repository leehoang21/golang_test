package models

import "base/internal/base/model"

type Token struct {
	model.BaseModel `bson:",inline"`
	UserID          string `json:"user_id" bson:"user_id"`
	Platform        string `json:"platform" bson:"platform"`
	OrgID           string `json:"ord_id" bson:"ord_id"`
	Revoke          bool   `json:"revoke" bson:"revoke"`
}
