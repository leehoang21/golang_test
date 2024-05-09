package querymgo

import (
	"base/internal/base/mgo/db"
	"base/internal/models"
	"base/internal/repository"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type tokenRepository struct {
	*db.Table
}

func NewTokenRepo(dbc *mongo.Database, name, prefix string) repository.Token {
	return tokenRepository{db.NewTable(name, prefix, dbc)}
}

func (u tokenRepository) RevokeAllByUserID(ctx context.Context, userID string) error {
	_, err := u.DeleteMany(ctx, bson.M{
		"user_id": userID,
	})
	return err
}

func (u tokenRepository) RevokeID(ctx context.Context, id string) error {
	return u.R_DeleteByID(ctx, id)
}

func (u tokenRepository) GetByID(ctx context.Context, id string) (*models.Token, error) {
	var usr *models.Token
	return usr, u.R_SelectByID(ctx, id, &usr)
}
