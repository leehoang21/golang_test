package repository

import (
	"base/internal/base/mgo"
	"base/internal/models"
	"context"
)

type GroupRole interface {
	mgo.BaseTable
	GetByName(ctx context.Context, name string) (*models.GroupRole, error)
	GetByUserID(ctx context.Context, userID string) ([]models.GroupRole, error)
	GetByID(ctx context.Context, id string) (*models.GroupRole, error)
}
