package repository

import (
	"base/internal/base/mgo"
	"base/internal/models"
	"context"
)

type User interface {
	mgo.BaseTable
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetByPhone(ctx context.Context, phone string) (*models.User, error)
	GetByID(ctx context.Context, id string) (*models.User, error)
}
