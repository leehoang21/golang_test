package models

import (
	"base/internal/base/model"
	"time"
)

type Estimate struct {
	model.BaseModel `bson:",inline"`
	TaskID          string    `json:"task_id" bson:"task_id"`
	EstimateTime    uint      `json:"estimate_time" bson:"estimate_time"`
	ActualTime      uint      `json:"actual_time" bson:"actual_time"`
	DateComplete    time.Time `json:"date_complete" bson:"date_complete"`
	CreatedByEmail  string    `json:"create_by_email" bson:"create_by_email"`
	CreateBy        *User     `json:"create_by" `
}
