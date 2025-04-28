package token

import (
	"base/internal/models"
)

type CreateInput struct {
	UserID   string
	Platform string
	OrgID    string
}

func (t CreateInput) ToModel() *models.Token {
	return &models.Token{
		UserID:   t.UserID,
		Platform: t.Platform,
		OrgID:    t.Platform,
	}
}
