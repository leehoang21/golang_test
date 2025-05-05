package models

import (
	"base/internal/base/model"
	"time"
)

type Task struct {
	model.BaseModel `bson:",inline"`
	Title           string    `json:"title" bson:"title"`
	Description     string    `json:"description" bson:"description"`
	AssignedToEmail string    `json:"assigned_to_email" bson:"assigned_to_email"`
	AssignedTo      *User     `json:"assigned_to" `
	Status          string    `json:"status" bson:"status"`
	Deadline        time.Time `json:"deadline" bson:"deadline"`
	CreatedByEmail  string    `json:"created_by_email" bson:"created_by_email"`
	CreatedBy       *User     `json:"created_by" `
	Estimate        *Estimate `json:"estimate" `
}
