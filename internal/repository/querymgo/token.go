package querymgo

import (
	"context"
	"eclectric/internal/base/mgo/db"
	"eclectric/internal/models"
	"eclectric/internal/repository"

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
	return u.UpdateAll(ctx, bson.M{
		"user_id": userID,
	}, bson.M{
		"revoke": true,
	})
}

func (u tokenRepository) RevokeID(ctx context.Context, id string) error {
	return u.UpdateAll(ctx, bson.M{
		"_id": id,
	}, bson.M{
		"revoke": true,
	})
}

func (u tokenRepository) GetByID(ctx context.Context, id string) (*models.Token, error) {
	var usr *models.Token
	return usr, u.SelectByID(ctx, id, &usr)
}
