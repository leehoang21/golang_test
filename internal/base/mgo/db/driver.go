package db

import (
	"context"
	"reflect"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type Infrastructure struct {
	m      sync.Mutex
	mongos map[string]*mongo.Client
}

func NewInfrastructure() *Infrastructure {
	return &Infrastructure{
		mongos: map[string]*mongo.Client{},
	}
}

func (inf *Infrastructure) ConnectMongo(ctx context.Context, uri, user, pass string) (client *mongo.Client, err error) {
	inf.m.Lock()
	defer inf.m.Unlock()
	client, ok := inf.mongos[uri]
	if ok {
		return
	}
	rb := bson.NewRegistryBuilder()
	rb.RegisterTypeMapEntry(bsontype.EmbeddedDocument, reflect.TypeOf(bson.M{}))
	var optionsClient = options.Client().ApplyURI(uri).SetRegistry(rb.Build())

	if user != "" && pass != "" {
		optionsClient.SetAuth(options.Credential{
			Username: user,
			Password: pass,
		})
	}

	client, err = mongo.NewClient(optionsClient)
	if err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return
	}
	inf.mongos[uri] = client
	return client, err
}

func (inf *Infrastructure) GetDatabase(client *mongo.Client, dbName string) *mongo.Database {
	return client.Database(dbName)
}

func R_Transaction(ctx context.Context, client *mongo.Client, mFuncs []func() error) error {
	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)

	session, err := client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	err = mongo.WithSession(ctx, session, func(sessionContext mongo.SessionContext) error {
		if err = session.StartTransaction(txnOpts); err != nil {
			return err
		}
		//start function create/insert/update
		for _, itemFunc := range mFuncs {
			err = itemFunc()
			if err != nil {
				return err
			}
		}
		//end funcions
		if err = session.CommitTransaction(sessionContext); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		if abortErr := session.AbortTransaction(ctx); abortErr != nil {
			return abortErr
		}
		return err
	}
	return nil
}
