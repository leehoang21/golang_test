package group_role

import (
	"base/internal/data/enums"
	"base/internal/models"
)

type GroupRoleInput struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Status      string   `json:"status" validate:"required"`
	Members     []string `json:"members" validate:"required"`
}

func (t GroupRoleInput) ToModel() *models.GroupRole {
	return &models.GroupRole{
		Name:        t.Name,
		Description: t.Description,
		Status:      enums.StringToStatusType(t.Status),
		MemberIds:   t.Members,
	}
}
