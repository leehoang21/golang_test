package querymgo

import (
	"base/internal/base/mgo/db"
	"base/internal/models"
	"base/internal/repository"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type estimateRepository struct {
	*db.Table
}

func NewEstimateRepo(dbc *mongo.Database, name, prefix string) repository.Estimate {
	return estimateRepository{db.NewTable(name, prefix, dbc)}
}

func (u estimateRepository) GetByID(ctx context.Context, id string) (*models.Estimate, error) {
	var estimate *models.Estimate
	return estimate, u.R_SelectByID(ctx, id, &estimate)
}

func (u estimateRepository) GetByTaskID(ctx context.Context, id string) (*models.Estimate, error) {
	var estimate *models.Estimate
	return estimate, u.R_SelectOne(ctx, bson.M{
		"task_id": id,
	}, &estimate)
}
