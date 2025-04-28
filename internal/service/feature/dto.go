package feature

import (
	"base/internal/data/enums"
	"base/internal/models"
)

type FeatureInput struct {
	Name      string   `json:"name" validate:"required"`
	KeyMenu   string   `json:"key_menu"`
	URLView   string   `json:"url_view" `
	Api       string   `json:"api" validate:"required"`
	Action    []string `json:"action" validate:"required"`
	Status    string   `json:"status" validate:"required"`
	RoleNames []string `json:"role_names" bson:"role_names" validate:"required"`
}

func (f FeatureInput) ToModel() *models.Feature {
	return &models.Feature{
		Name:      f.Name,
		KeyMenu:   f.KeyMenu,
		URLView:   f.URLView,
		Api:       f.Api,
		Action:    f.Action,
		Status:    enums.StringToStatusType(f.Status),
		RoleNames: f.RoleNames,
	}
}
