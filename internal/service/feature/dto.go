package feature

import (
	"base/internal/models"
)

type Input struct {
	Name      string   `json:"name" validate:"required"`
	KeyMenu   string   `json:"key_menu"`
	URLView   string   `json:"url_view" `
	Api       string   `json:"api" validate:"required"`
	Action    []string `json:"action" validate:"required"`
	Status    string   `json:"status" validate:"required"`
	RoleNames []string `json:"role_names" bson:"role_names"`
}

func (i Input) ToModel() *models.Feature {
	return &models.Feature{
		Name:      i.Name,
		KeyMenu:   i.KeyMenu,
		URLView:   i.URLView,
		Api:       i.Api,
		Action:    i.Action,
		Status:    i.Status,
		RoleNames: i.RoleNames,
	}

}
