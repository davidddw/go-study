package main

import (
	"context"
	"fmt"
	"runtime/debug"
	"sync"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// BaseEntity base
type BaseEntity interface {
	GetID() string
	SetID(id string)
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
	host = "localhost"
	user = "cloud"
	pass = "passwd"
	db   = "mydb"
)

func init() {
	SetConnect(host, user, pass, db)
}

// SetConnect init connect
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

// Insert insert one record to DB
func (m *MongoClient) Insert(collection string, e BaseEntity) (s string, err error) {
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
	e.SetID(uuid.New().String())
	cid, err := collections.InsertOne(m.Ctx, e)
	if err != nil {
		return
	}
	s = cid.InsertedID.(primitive.ObjectID).Hex()
	return
}

// GetOneByID select one record from DB
func (m *MongoClient) GetOneByID(collection string, id string, e BaseEntity) (err error) {
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
	result := collections.FindOne(m.Ctx, bson.M{"id": id})
	result.Decode(e)
	return
}

// Count get count
func (m *MongoClient) Count(collection string, filter PageFilter) (c int64, err error) {
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

// GetAllByFilter get all
func (m *MongoClient) GetAllByFilter(collection string, filter PageFilter, e *[]*Book) (err error) {
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
	err = cur.All(m.Ctx, e)
	fmt.Printf("%v\n", e)
	// for cur.Next(m.Ctx) {
	// 	var result BaseEntity
	// 	cur.Decode(&result)
	// 	e = append(e, result)
	// 	//var e interface{}
	// 	//cur.Decode(&e)
	// }
	return
}
