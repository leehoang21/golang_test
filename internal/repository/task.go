package repository

import (
	"base/internal/base/mgo"
	"base/internal/models"
	"context"
)

type Task interface {
	mgo.BaseTable
	GetByID(ctx context.Context, id string) (*models.Task, error)
	GetByTitle(ctx context.Context, title string) (*models.Task, error)
}
