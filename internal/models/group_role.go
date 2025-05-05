package models

import (
	"base/internal/base/model"
)

type GroupRole struct {
	model.BaseModel `bson:",inline"`
	Name            string   `json:"name" bson:"name"`
	Description     string   `json:"description" bson:"description"`
	Status          string   `json:"status" bson:"status"`
	MemberIds       []string `json:"member_ids" bson:"member_ids"`
	Members         []User   `json:"members"`
}
