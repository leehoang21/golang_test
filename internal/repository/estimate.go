package repository

import (
	"base/internal/base/mgo"
	"base/internal/models"
	"context"
)

type Estimate interface {
	mgo.BaseTable
	GetByID(ctx context.Context, id string) (*models.Estimate, error)
	GetByTaskID(ctx context.Context, id string) (*models.Estimate, error)
}
