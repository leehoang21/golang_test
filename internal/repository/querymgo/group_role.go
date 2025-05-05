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

type groupRoleRepository struct {
	*db.Table
}

func NewRoleRepo(dbc *mongo.Database, name, prefix string) repository.GroupRole {
	return groupRoleRepository{db.NewTable(name, prefix, dbc)}
}

func (u groupRoleRepository) GetByName(ctx context.Context, name string) (*models.GroupRole, error) {
	var role *models.GroupRole
	return role, u.R_SelectOne(ctx, bson.M{
		"name": name,
	}, &role)
}

func (u groupRoleRepository) GetByID(ctx context.Context, id string) (*models.GroupRole, error) {
	var role *models.GroupRole

	pipeline := mongo.Pipeline{
		{{"$match", bson.D{
			{"_id", id},
			{"dtime", 0},
		}}},
		{{"$limit", 1}},
		{{"$lookup", bson.D{
			{"from", "users"},
			{"let", bson.D{{"member_ids", "$roles.member_ids"}}},
			{"pipeline", bson.A{
				bson.D{{"$match", bson.D{
					{"$expr", bson.D{
						{"$and", bson.A{
							bson.D{{"$in", bson.A{"$_id", "$$member_ids"}}},
							bson.D{{"$eq", bson.A{"$dtime", 0}}},
						}},
					}},
				}}},
			}},
			{"as", "members"},
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
		if err := showInfoCursor.Decode(&role); err != nil {
			return nil, fmt.Errorf("decode error: %w", err)
		}
		return role, nil
	}

	return nil, err
}

func (u groupRoleRepository) GetByUserID(ctx context.Context, userID string) ([]models.GroupRole, error) {
	var roles []models.GroupRole
	return roles, u.R_SelectMany(ctx, bson.M{
		"member_ids": userID,
	}, &roles)
}
