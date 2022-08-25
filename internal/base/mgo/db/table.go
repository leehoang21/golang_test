package db

import (
	"context"
	"fmt"
	"time"

	"eclectric/internal/base/mgo/filter"
	"eclectric/internal/base/mlog"
	"eclectric/internal/base/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var logDB = mlog.NewTagLog("MONGO_DB")

type Table struct {
	*mongo.Collection
	Prefix string
}

func NewTable(name, prefix string, db *mongo.Database) *Table {
	fmt.Println("DB", name)
	return &Table{
		Collection: db.Collection(name),
		Prefix:     prefix,
	}
}

func (t *Table) Search(ctx context.Context, f filter.Filter, val interface{}) error {
	var q = []bson.M{}
	//var opts = options.Find()
	if f.GetWhere() != nil {
		q = append(q, bson.M{"$match": f.GetWhere()})
	}

	if f.GetOrderBy() != nil {
		q = append(q, bson.M{"$sort": f.GetOrderBy()})
	}

	if f.GetOffset() > 0 {
		//opts.SetSkip(f.GetOffset())
		q = append(q, bson.M{"$skip": f.GetOffset()})
	}
	if f.GetLimit() > 0 {
		//opts.SetLimit(f.GetLimit())
		q = append(q, bson.M{"$limit": f.GetLimit()})
	}

	if f.GetJoins() != nil {
		q = append(q, bson.M{"$lookup": f.GetJoins()})
	}

	return t.Pipe(ctx, q, val)
}

func (t *Table) CreateIndexOne(ctx context.Context, mode mongo.IndexModel, opts ...*options.CreateIndexesOptions) error {
	var _, err = t.Indexes().CreateOne(ctx, mode, opts...)
	return err
}

func (t *Table) CreateIndexMany(ctx context.Context, mods []mongo.IndexModel, opts ...*options.CreateIndexesOptions) error {

	var _, err = t.Indexes().CreateMany(ctx, mods, opts...)
	return err
}

func (t *Table) Create(ctx context.Context, model model.IModel) error {

	model.BeforeCreate(t.Prefix)
	var _, err = t.InsertOne(ctx, model)
	if err != nil {
		logDB.Errorf("Create table "+t.Name()+": "+err.Error(), model)
	}
	return err
}

func (t *Table) CreateForce(ctx context.Context, model model.IModel) error {

	var _, err = t.InsertOne(ctx, model)
	if err != nil {
		logDB.Errorf("Create table "+t.Name()+": "+err.Error(), model)
	}
	return err
}

func (t *Table) Update(ctx context.Context, model model.IModel) error {
	model.BeforeUpdate()
	var _, err = t.UpdateOne(ctx, bson.M{"_id": model.GetID(), "dtime": 0}, bson.M{"$set": model})
	if err != nil {
		logDB.Errorf("Update table "+t.Name()+": "+err.Error(), model)
	}
	return err
}

func (t *Table) Delete(ctx context.Context, id string, model model.IModel) error {
	model.BeforeDelete()
	var _, err = t.UpdateOne(ctx, bson.M{"_id": id, "dtime": 0}, bson.M{"$set": bson.M{"dtime": time.Now().Unix()}})
	if err != nil {
		logDB.Errorf("Delete table "+t.Name()+": "+err.Error(), model)
	}
	return err
}

func (t *Table) DeleteByID(ctx context.Context, id string) error {
	var _, err = t.UpdateByID(ctx, id, bson.M{"$set": bson.M{"dtime": time.Now().Unix()}})
	if err != nil {
		logDB.Errorf("DeleteByID "+err.Error(), id)
	}
	return err
}
func (t *Table) SelectAndDelete(ctx context.Context, id string) error {
	var timeNow = time.Now().Unix()
	after := options.After
	opts := &options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	var res = t.FindOneAndUpdate(ctx, bson.M{"_id": id, "dtime": 0},
		bson.M{"$set": bson.M{"dtime": timeNow, "ctime": timeNow}}, opts)
	if res.Err() != nil {
		logDB.Errorf("Delete table " + t.Name() + ": " + res.Err().Error())
	}
	return res.Err()
}

func (t *Table) UnsafeUpdate(ctx context.Context, filter bson.M, v interface{}) error {
	filter["dtime"] = 0
	var _, err = t.UpdateOne(ctx, filter,
		bson.M{"$set": v})
	if err != nil {
		logDB.Errorf("UnsafeUpdate table "+t.Name()+": "+err.Error(), v)
	}
	return err
}

func (t *Table) UpdateForce(ctx context.Context, filter bson.M, v interface{}) error {
	filter["dtime"] = 0
	var _, err = t.UpdateOne(ctx, filter, v)
	if err != nil {
		logDB.Errorf("UnsafeUpdate table "+t.Name()+": "+err.Error(), v)
	}
	return err
}

func (t *Table) UnsafeUpdateByID(ctx context.Context, id string, v interface{}) error {
	var _, err = t.UpdateOne(ctx,
		bson.M{"dtime": 0,
			"_id": id},
		bson.M{"$set": v})
	if err != nil {
		logDB.Errorf("UnsafeUpdateByID table "+t.Name()+": "+err.Error(), v)
	}
	return err
}

func (t *Table) CreateMany(ctx context.Context, v []interface{}) ([]interface{}, error) {
	var res, err = t.InsertMany(ctx, v)
	var ids []interface{}
	if err != nil {
		logDB.Errorf("UnsafeUpdateByID table "+t.Name()+": "+err.Error(), v)
	}
	if res != nil {
		ids = res.InsertedIDs
	}
	return ids, err
}

func (t *Table) SelectOne(ctx context.Context, filter bson.M, v interface{}) error {
	filter["dtime"] = 0
	var err = t.FindOne(ctx, filter).Decode(v)
	return err
}

func (t *Table) SelectOneWithFields(ctx context.Context, filter bson.M, v interface{}, fields bson.M) error {
	filter["dtime"] = 0
	var opts = options.FindOne().SetProjection(fields)
	var err = t.FindOne(ctx, filter, opts).Decode(v)
	return err
}

func (t *Table) SelectManyWithFields(ctx context.Context, filter bson.M, v interface{}, fields bson.M) error {
	filter["dtime"] = 0
	var opts = options.Find().SetProjection(fields)
	var cur, err = t.Find(ctx, filter, opts)
	if err != nil {
		cur.Close(ctx)
		return err
	}
	err = cur.All(ctx, v)
	return err
}

func (t *Table) SelectByID(ctx context.Context, id string, v interface{}) error {
	var filter = bson.M{
		"dtime": 0,
		"_id":   id,
	}
	var err = t.FindOne(ctx, filter).Decode(v)

	return err
}

func (t *Table) SelectMany(ctx context.Context, filter bson.M, v interface{}) error {

	filter["dtime"] = 0
	var cur, err = t.Find(ctx, filter)
	if err != nil {
		cur.Close(ctx)
		return err
	}
	err = cur.All(ctx, v)
	return err
}

func (t *Table) SelectDistinct(ctx context.Context, field string, filter bson.M) ([]interface{}, error) {

	filter["dtime"] = 0
	return t.Distinct(ctx, field, filter)
}

func (t *Table) UpdateAll(ctx context.Context, filter bson.M, update interface{}) error {

	filter["dtime"] = 0
	var _, err = t.UpdateMany(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return err
	}
	return err
}

func (t *Table) SelectAndSort(ctx context.Context, filter bson.M, sortFields bson.M, skip, limit int64, res interface{}) error {

	filter["dtime"] = 0
	var opts = options.Find()
	if sortFields != nil {
		// sort := bson.M{}
		// for key, val := range sortFields {
		// 	sort = append(sort, bson.E{key: val})
		// }
		opts.SetSort(sortFields)
	}
	if skip > 0 {
		opts.SetSkip(skip)
	}
	if limit > 0 {
		opts.SetLimit(limit)
	}
	var cur, err = t.Find(ctx, filter, opts)
	if err != nil {
		return err
	}
	err = cur.All(ctx, res)
	return err
}

func (t *Table) Pipe(ctx context.Context, pipeline []bson.M, res interface{}) error {

	var cur, err = t.Aggregate(ctx, pipeline)
	if err != nil {
		return err
	}
	err = cur.All(ctx, res)
	return err
}

func (t *Table) Count(ctx context.Context, filter bson.M) (int64, error) {
	filter["dtime"] = 0

	return t.CountDocuments(ctx, filter)
}
