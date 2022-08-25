package querymgo

import (
	"context"
	"eclectric/internal/base/mgo/db"
	"eclectric/internal/models"
	"eclectric/internal/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	*db.Table
}

func NewUserRepo(dbc *mongo.Database, name, prefix string) repository.User {
	return userRepository{db.NewTable(name, prefix, dbc)}
}

func (u userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var usr *models.User
	return usr, u.SelectOne(ctx, bson.M{
		"email": email,
	}, &usr)
}

func (u userRepository) GetByPhone(ctx context.Context, phone string) (*models.User, error) {
	var usr *models.User
	return usr, u.SelectOne(ctx, bson.M{
		"phone": phone,
	}, &usr)
}

func (u userRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	var usr *models.User
	return usr, u.SelectByID(ctx, id, &usr)
}
