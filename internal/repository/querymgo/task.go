package querymgo

import (
	"base/internal/base/mgo/db"
	"base/internal/models"
	"base/internal/repository"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type taskRepository struct {
	*db.Table
}

func NewTaskRepo(dbc *mongo.Database, name, prefix string) repository.Task {
	return taskRepository{db.NewTable(name, prefix, dbc)}
}

func (u taskRepository) GetByID(ctx context.Context, id string) (*models.Task, error) {
	var task *models.Task

	pipeline := mongo.Pipeline{
		{{"$match", bson.D{
			{"_id", id},
			{"dtime", 0}}}},
		{{"$limit", 1}},
		{{"$lookup", bson.D{
			{"from", "users"},
			{"let", bson.D{{"email", "$assigned_to_email"}}},
			{"pipeline", bson.A{
				bson.D{{"$match", bson.D{
					{"$expr", bson.D{
						{"$and", bson.A{
							bson.D{{"$eq", bson.A{"$email", "$$email"}}},
							bson.D{{"$eq", bson.A{"$dtime", 0}}},
						}},
					}},
				}}},
			}},
			{"as", "assigned_to"},
		}}},
		{{"$unwind", bson.D{
			{"path", "$assigned_to"},
			{"preserveNullAndEmptyArrays", true},
		}}},

		{{"$lookup", bson.D{
			{"from", "users"},
			{"let", bson.D{{"email", "$created_by_email"}}},
			{"pipeline", bson.A{
				bson.D{{"$match", bson.D{
					{"$expr", bson.D{
						{"$and", bson.A{
							bson.D{{"$eq", bson.A{"$email", "$$email"}}},
							bson.D{{"$eq", bson.A{"$dtime", 0}}},
						}},
					}},
				}}},
			}},
			{"as", "created_by"},
		}}},
		{{"$unwind", bson.D{
			{"path", "$created_by"},
			{"preserveNullAndEmptyArrays", true},
		}}},

		{{"$lookup", bson.D{
			{"from", "estimates"},
			{"let", bson.D{{"id", "$_id"}}},
			{"pipeline", bson.A{
				bson.D{{"$match", bson.D{
					{"$expr", bson.D{
						{"$and", bson.A{
							bson.D{{"$eq", bson.A{"$task_id", "$$id"}}},
							bson.D{{"$eq", bson.A{"$dtime", 0}}},
						}},
					}},
				}}},
			}},
			{"as", "estimate"},
		}}},
		{{"$unwind", bson.D{
			{"path", "$estimate"},
			{"preserveNullAndEmptyArrays", true},
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
		if err := showInfoCursor.Decode(&task); err != nil {
			return nil, fmt.Errorf("decode error: %w", err)
		}
		return task, nil
	}

	return nil, err
}

func (u taskRepository) GetByTitle(ctx context.Context, title string) (*models.Task, error) {
	var task *models.Task
	return task, u.R_SelectOne(ctx, bson.M{
		"title": title,
	}, &task)
}
