package repository

import (
	"base/internal/base/mgo"
	"base/internal/models"
	"context"
)

type Feature interface {
	mgo.BaseTable
	GetByName(ctx context.Context, name string) (*models.Feature, error)
	GetByApi(ctx context.Context, api string) (*models.Feature, error)
	GetByID(ctx context.Context, id string) (*models.Feature, error)
}
