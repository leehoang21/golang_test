package querymgo

import (
	"base/internal/base/mgo/db"
	"base/internal/models"
	"base/internal/repository"
	"context"

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
	err := u.R_SelectOne(ctx, bson.M{
		"email": email,
	}, &usr)
	return usr, err

}

func (u userRepository) GetByPhone(ctx context.Context, phone string) (*models.User, error) {
	var usr *models.User
	return usr, u.R_SelectOne(ctx, bson.M{
		"phone": phone,
	}, &usr)
}

func (u userRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	var usr *models.User
	return usr, u.R_SelectByID(ctx, id, &usr)
}
