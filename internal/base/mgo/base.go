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
	R_Search(ctx context.Context, f filter.Filter, val interface{}) error
	R_SearchAndCount(ctx context.Context, f filter.Filter, val interface{}) (int64, error)
	R_CreateIndexOne(ctx context.Context, mode mongo.IndexModel, opts ...*options.CreateIndexesOptions) error
	R_CreateIndexMany(ctx context.Context, mods []mongo.IndexModel, opts ...*options.CreateIndexesOptions) error
	R_Create(ctx context.Context, model model.IModel) error
	R_CreateForce(ctx context.Context, model model.IModel) error
	R_Update(ctx context.Context, model model.IModel) error
	R_Delete(ctx context.Context, id string, model model.IModel) error
	R_DeleteByID(ctx context.Context, id string) error
	R_SelectAndDelete(ctx context.Context, id string) error
	R_UnsafeUpdate(ctx context.Context, filter bson.M, v interface{}) error
	R_UpdateForce(ctx context.Context, filter bson.M, v interface{}) error
	R_UnsafeUpdateByID(ctx context.Context, id string, v interface{}) error
	R_CreateMany(ctx context.Context, v []interface{}) ([]interface{}, error)
	R_SelectOne(ctx context.Context, filter bson.M, v interface{}) error
	R_SelectOneWithFields(ctx context.Context, filter bson.M, v interface{}, fields bson.M) error
	R_SelectManyWithFields(ctx context.Context, filter bson.M, v interface{}, fields bson.M) error
	R_SelectByID(ctx context.Context, id string, v interface{}) error
	R_SelectMany(ctx context.Context, filter bson.M, v interface{}) error
	R_SelectDistinct(ctx context.Context, field string, filter bson.M) ([]interface{}, error)
	R_UpdateAll(ctx context.Context, filter bson.M, update interface{}) error
	R_SelectAndSort(ctx context.Context, filter bson.M, sortFields bson.M, skip, limit int64, res interface{}) error
	R_Pipe(ctx context.Context, pipeline []bson.M, res interface{}) error
	R_Count(ctx context.Context, filter bson.M) (int64, error)
}

func NewBaseTable(name, prefix string, dbc *mongo.Database) BaseTable {
	return db.NewTable(name, prefix, dbc)
}
