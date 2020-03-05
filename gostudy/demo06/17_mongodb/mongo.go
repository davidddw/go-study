package main

import (
	"context"
	"fmt"
	"runtime/debug"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BaseEntity interface {
}

type MongoClient struct {
	Client   *mongo.Client
	Ctx      context.Context
	Database string
}

type PageFilter struct {
	SortBy     string
	SortMode   int8
	Limit      *int64
	Skip       *int64
	Filter     map[string]interface{}
	RegexFiler map[string]string
}

var Mongo *MongoClient
var cancel context.CancelFunc

const (
	host = "192.168.1.62"
	user = "cloud"
	pass = "passwd"
	db   = "mydb"
)

func init() {
	SetConnect(host, user, pass, db)
}

func SetConnect(host, user, pass, db string) {
	var once sync.Once
	once.Do(func() {
		var ctx context.Context
		ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
		opts := &options.ClientOptions{}
		opts.SetAuth(options.Credential{
			AuthMechanism: "SCRAM-SHA-1",
			AuthSource:    db,
			Username:      user,
			Password:      pass,
		}).ApplyURI(fmt.Sprintf("mongodb://%s:27017", host)).SetMaxPoolSize(20)
		client, err := mongo.Connect(ctx, opts)
		if err != nil {
			fmt.Println(err)

		}
		Mongo = &MongoClient{Ctx: ctx, Client: client, Database: db}
	})
}

func (m *MongoClient) Create(collection string, e BaseEntity) (error, string) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				debug.PrintStack()
			}
		}
	}()
	collections := m.Client.Database(m.Database).Collection(collection)
	cid, err := collections.InsertOne(m.Ctx, e)
	if err != nil {
		return err, ""
	}
	return nil, cid.InsertedID.(primitive.ObjectID).Hex()
}

func (m *MongoClient) Get(collection string, id string) (err error, e interface{}) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				debug.PrintStack()
			}
		}
	}()
	collections := m.Client.Database(m.Database).Collection(collection)
	objID, _ := primitive.ObjectIDFromHex(id)
	result := collections.FindOne(m.Ctx, bson.M{"_id": objID})
	result.Decode(&e)
	return
}

func (m *MongoClient) Count(collection string, filter PageFilter) (err error, c int64) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				debug.PrintStack()
			}
		}
	}()
	if filter.RegexFiler != nil {
		for k, v := range filter.RegexFiler {
			filter.Filter[k] = primitive.Regex{Pattern: v, Options: ""}
		}
	}
	collections := m.Client.Database(m.Database).Collection(collection)
	c, err = collections.CountDocuments(m.Ctx, filter.Filter, &options.CountOptions{Skip: filter.Skip, Limit: filter.Limit})
	return
}

func (m *MongoClient) List(collection string, filter PageFilter) (err error, e []interface{}) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				debug.PrintStack()
			}
		}
	}()
	if filter.RegexFiler != nil {
		for k, v := range filter.RegexFiler {
			filter.Filter[k] = primitive.Regex{Pattern: v, Options: ""}
		}
	}
	collections := m.Client.Database(m.Database).Collection(collection)
	cur, err := collections.Find(m.Ctx, filter.Filter, &options.FindOptions{Limit: filter.Limit, Skip: filter.Skip, Sort: bson.M{filter.SortBy: filter.SortMode}})
	//cur, err := collections.Find(m.Ctx, bson.D{})
	defer cur.Close(m.Ctx)
	for cur.Next(m.Ctx) {
		var result Book
		cur.Decode(&result)
		e = append(e, result)
		//var e interface{}
		//cur.Decode(&e)
	}
	return
}
