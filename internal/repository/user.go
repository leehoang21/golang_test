package repository

import (
	"context"
	"eclectric/internal/base/mgo"
	"eclectric/internal/models"
)

type User interface {
	mgo.BaseTable
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetByPhone(ctx context.Context, phone string) (*models.User, error)
}
