package task

import (
	"base/internal/models"
	"time"
)

type Input struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	AssignedTo  string    `json:"assigned_to" validate:"required"`
	Status      string    `json:"status"  validate:"required"`
	Deadline    time.Time `json:"deadline" validate:"required"`
}

func (i Input) ToModel() *models.Task {
	return &models.Task{
		Title:           i.Title,
		Description:     i.Description,
		AssignedToEmail: i.AssignedTo,
		Status:          i.Status,
		Deadline:        i.Deadline,
	}
}
