package group_role

import (
	"base/internal/models"
)

type GroupRoleInput struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Status      string   `json:"status" validate:"required"`
	MemberIds   []string `json:"member_ids"`
}

func (t GroupRoleInput) ToModel() *models.GroupRole {
	return &models.GroupRole{
		Name:        t.Name,
		Description: t.Description,
		Status:      t.Status,
		MemberIds:   t.MemberIds,
	}
}
