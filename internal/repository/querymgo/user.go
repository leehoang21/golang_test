package querymgo

import (
	"base/internal/base/mgo/db"
	"base/internal/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	*db.Table
}

func NewUserRepo(dbc *mongo.Database, name, prefix string) repository.User {
	return userRepository{db.NewTable(name, prefix, dbc)}
}
