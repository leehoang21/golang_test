package repository

import (
	"context"
	"eclectric/internal/base/mgo"
	"eclectric/internal/models"
)

type Token interface {
	mgo.BaseTable
	RevokeAllByUserID(ctx context.Context, userID string) error
	RevokeID(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*models.Token, error)
}
