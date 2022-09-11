package token

import (
	"base/internal/models"
)

type TokenCreateInput struct {
	UserID   string
	Platform string
	OrgID    string
}

func (t TokenCreateInput) ToModel() *models.Token {
	return &models.Token{
		UserID:   t.UserID,
		Platform: t.Platform,
		OrgID:    t.Platform,
	}
}
