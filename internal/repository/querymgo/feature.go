package querymgo

import (
	"base/internal/base/mgo/db"
	"base/internal/models"
	"base/internal/repository"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type featureRepository struct {
	*db.Table
}

func NewFeatureRepo(dbc *mongo.Database, name, prefix string) repository.Feature {
	return featureRepository{db.NewTable(name, prefix, dbc)}
}

func (u featureRepository) GetByName(ctx context.Context, name string) (*models.Feature, error) {
	var f *models.Feature
	return f, u.R_SelectOne(ctx, bson.M{
		"name": name,
	}, &f)
}

func (u featureRepository) GetByApi(ctx context.Context, api string) (*models.Feature, error) {
	var f *models.Feature
	return f, u.R_SelectOne(ctx, bson.M{
		"api": api,
	}, &f)
}

func (u featureRepository) GetByID(ctx context.Context, id string) (*models.Feature, error) {
	var result *models.Feature

	pipeline := mongo.Pipeline{
		{{"$match", bson.D{
			{"_id", id},
			{"dtime", 0},
		}}},
		{{"$limit", 1}},
		{{"$lookup", bson.D{
			{"from", "group_roles"},
			{"let", bson.D{{"role_names", "$role_names"}}},
			{"pipeline", bson.A{
				bson.D{{"$match", bson.D{
					{"$expr", bson.D{
						{"$and", bson.A{
							bson.D{{"$in", bson.A{"$name", "$$role_names"}}},
							bson.D{{"$eq", bson.A{"$dtime", 0}}},
						}},
					}},
				}}},
			}},
			{"as", "roles"},
		}}},
		{{"$unwind", bson.D{
			{"path", "$roles"},
			{"preserveNullAndEmptyArrays", true},
		}}},
		{{"$group", bson.D{
			{"_id", "$_id"},
			{"ctime", bson.D{{"$first", "$ctime"}}},
			{"mtime", bson.D{{"$first", "$mtime"}}},
			{"dtime", bson.D{{"$first", "$dtime"}}},
			{"name", bson.D{{"$first", "$name"}}},
			{"key_menu", bson.D{{"$first", "$key_menu"}}},
			{"url_view", bson.D{{"$first", "$url_view"}}},
			{"api", bson.D{{"$first", "$api"}}},
			{"action", bson.D{{"$first", "$action"}}},
			{"status", bson.D{{"$first", "$status"}}},
			{"roles", bson.D{{"$push", "$roles"}}},
		}}},
	}

	showInfoCursor, err := u.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer func(showInfoCursor *mongo.Cursor, ctx context.Context) {
		err := showInfoCursor.Close(ctx)
		if err != nil {
			log.Fatal("Error closing cursor", err)
		}
	}(showInfoCursor, context.TODO())
	if showInfoCursor.Next(ctx) {
		if err := showInfoCursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("decode error: %w", err)
		}
		return result, nil
	}

	return nil, err
}
