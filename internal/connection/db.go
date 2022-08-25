package connection

import (
	"context"
	"eclectric/config"
	mgo "eclectric/internal/base/mgo/db"

	"go.mongodb.org/mongo-driver/mongo"
)

func ConnectDB(ctx context.Context, dbConfig config.DB) *mongo.Database {
	var intra = mgo.NewInfrastructure()
	client, err := intra.ConnectMongo(ctx, dbConfig.Path, dbConfig.DBUser, dbConfig.DBPass)
	if err != nil {
		panic(err)
	}
	return intra.GetDatabase(client, dbConfig.DBName)
}
