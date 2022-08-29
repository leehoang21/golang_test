package mgo

import (
	"base/internal/base/mgo/db"
	"base/internal/base/mgo/filter"
	"base/internal/base/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BaseTable interface {
	Search(ctx context.Context, f filter.Filter, val interface{}) error
	CreateIndexOne(ctx context.Context, mode mongo.IndexModel, opts ...*options.CreateIndexesOptions) error
	CreateIndexMany(ctx context.Context, mods []mongo.IndexModel, opts ...*options.CreateIndexesOptions) error
	Create(ctx context.Context, model model.IModel) error
	CreateForce(ctx context.Context, model model.IModel) error
	Update(ctx context.Context, model model.IModel) error
	Delete(ctx context.Context, id string, model model.IModel) error
	DeleteByID(ctx context.Context, id string) error
	SelectAndDelete(ctx context.Context, id string) error
	UnsafeUpdate(ctx context.Context, filter bson.M, v interface{}) error
	UpdateForce(ctx context.Context, filter bson.M, v interface{}) error
	UnsafeUpdateByID(ctx context.Context, id string, v interface{}) error
	CreateMany(ctx context.Context, v []interface{}) ([]interface{}, error)
	SelectOne(ctx context.Context, filter bson.M, v interface{}) error
	SelectOneWithFields(ctx context.Context, filter bson.M, v interface{}, fields bson.M) error
	SelectManyWithFields(ctx context.Context, filter bson.M, v interface{}, fields bson.M) error
	SelectByID(ctx context.Context, id string, v interface{}) error
	SelectMany(ctx context.Context, filter bson.M, v interface{}) error
	SelectDistinct(ctx context.Context, field string, filter bson.M) ([]interface{}, error)
	UpdateAll(ctx context.Context, filter bson.M, update interface{}) error
	SelectAndSort(ctx context.Context, filter bson.M, sortFields bson.M, skip, limit int64, res interface{}) error
	Pipe(ctx context.Context, pipeline []bson.M, res interface{}) error
	Count(ctx context.Context, filter bson.M) (int64, error)
}

func NewBaseTable(name, prefix string, dbc *mongo.Database) BaseTable {
	return db.NewTable(name, prefix, dbc)
}
