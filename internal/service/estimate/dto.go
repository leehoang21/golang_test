package estimate

import (
	"base/internal/models"
	"time"
)

type Input struct {
	TaskID       string    `json:"task_id" validate:"required"`
	EstimateTime uint      `json:"estimate_time" validate:"required"`
	ActualTime   uint      `json:"actual_time"`
	DateComplete time.Time `json:"date_complete"`
}

func (i Input) ToModel() *models.Estimate {
	return &models.Estimate{
		TaskID:       i.TaskID,
		EstimateTime: i.EstimateTime,
		ActualTime:   i.ActualTime,
		DateComplete: i.DateComplete,
	}
}
