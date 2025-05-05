package models

import (
	"base/internal/base/model"
)

type Feature struct {
	model.BaseModel `bson:",inline"`
	Name            string      `json:"name" bson:"name"`
	KeyMenu         string      `json:"key_menu" bson:"key_menu"`
	URLView         string      `json:"url_view" bson:"url_view"`
	Api             string      `json:"api" bson:"api"`
	Action          []string    `json:"action" bson:"action"`
	Status          string      `json:"status" bson:"status"`
	RoleNames       []string    `json:"role_names" bson:"role_names"`
	Roles           []GroupRole `json:"roles"`
}
