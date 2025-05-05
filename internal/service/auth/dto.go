package auth

import (
	"base/internal/models"
)

type ResponseLogin struct {
	AccessToken string       `json:"access_token"`
	User        *models.User `json:"user"`
}
