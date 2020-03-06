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
	host = "192.168.1.62"
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

// RecoverDebug if error occur Recover env and PrintStack
func RecoverDebug() {
	if r := recover(); r != nil {
		if _, ok := r.(error); !ok {
			debug.PrintStack()
		}
	}
}

// Create insert one record to DB
func (m *MongoClient) Create(collection string, e BaseEntity) (s string, err error) {
	defer RecoverDebug()
	collections := m.Client.Database(m.Database).Collection(collection)
	cid, err := collections.InsertOne(m.Ctx, e)
	if err != nil {
		return
	}
	s = cid.InsertedID.(primitive.ObjectID).Hex()
	return
}

// Create insert one record to DB
func (m *MongoClient) CreateMany(collection string, all []interface{}) (s []string, err error) {
	defer RecoverDebug()
	collections := m.Client.Database(m.Database).Collection(collection)
	cid, err := collections.InsertMany(m.Ctx, all)
	if err != nil {
		return
	}
	for _, v := range cid.InsertedIDs {
		s = append(s, v.(primitive.ObjectID).Hex())
	}
	return
}

// GetOneByID select one record from DB by id
// id is hex string
// e is pointer to object
func (m *MongoClient) GetOneByID(collection string, id string, e BaseEntity) (err error) {
	defer RecoverDebug()
	collections := m.Client.Database(m.Database).Collection(collection)
	objID, _ := primitive.ObjectIDFromHex(id)
	singleResult := collections.FindOne(m.Ctx, bson.M{"_id": objID})
	singleResult.Decode(e)
	return
}

// GetOneByUUID select one record from DB by UUID
// id is uuid string, e is pointer to object
// var o Object
// GetOneByUUID("book", "82164f43-a04d-4e9f-9cf4-45dc710e3f9c", &o)
func (m *MongoClient) GetOneByUUID(collection string, id string, e BaseEntity) (err error) {
	defer RecoverDebug()
	collections := m.Client.Database(m.Database).Collection(collection)
	singleResult := collections.FindOne(m.Ctx, bson.M{"id": id})
	singleResult.Decode(e)
	return
}

// Count get count
func (m *MongoClient) Count(collection string, filter PageFilter) (c int64, err error) {
	defer RecoverDebug()
	if filter.RegexFiler != nil {
		for k, v := range filter.RegexFiler {
			filter.Filter[k] = primitive.Regex{Pattern: v, Options: ""}
		}
	}
	collections := m.Client.Database(m.Database).Collection(collection)
	c, err = collections.CountDocuments(m.Ctx, filter.Filter,
		&options.CountOptions{Skip: filter.Skip, Limit: filter.Limit})
	return
}

// ListByFilter get all by PageFilter
// all is pointer to object
//     var all []Object
// 	   Mongo.GetAllByFilter("ob", filter, &all)
func (m *MongoClient) ListByFilter(collection string, filter PageFilter, all interface{}) (err error) {
	defer RecoverDebug()
	if filter.RegexFiler != nil {
		for k, v := range filter.RegexFiler {
			filter.Filter[k] = primitive.Regex{Pattern: v, Options: ""}
		}
	}
	collections := m.Client.Database(m.Database).Collection(collection)
	cursor, err := collections.Find(m.Ctx, filter.Filter, &options.FindOptions{
		Limit: filter.Limit,
		Skip:  filter.Skip,
		Sort:  bson.M{filter.SortBy: filter.SortMode},
	})
	defer cursor.Close(m.Ctx)
	err = cursor.All(m.Ctx, all)
	return
}

// ModifyByID update data
func (m *MongoClient) ModifyByID(collection string, e BaseEntity) (int64, error) {
	defer RecoverDebug()
	collections := m.Client.Database(m.Database).Collection(collection)
	// collections.UpdateOne
	// collections.UpdateMany
	result, err := collections.ReplaceOne(m.Ctx, bson.M{"id": e.GetID()}, e)
	return result.ModifiedCount, err
}

// Delete delete
func (m *MongoClient) Delete(collection, id string) (int64, error) {
	defer RecoverDebug()
	collections := m.Client.Database(m.Database).Collection(collection)
	objID, _ := primitive.ObjectIDFromHex(id)
	result, err := collections.DeleteOne(m.Ctx, bson.M{"_id": objID})
	return result.DeletedCount, err
}

// DeleteManyByRegex delete
func (m *MongoClient) DeleteMany(collection string, key string, value interface{}) (int64, error) {
	defer RecoverDebug()
	collections := m.Client.Database(m.Database).Collection(collection)
	filter := bson.D{primitive.E{Key: key, Value: value}}
	result, err := collections.DeleteMany(m.Ctx, filter)
	return result.DeletedCount, err
}

func UUID() string {
	return uuid.New().String()
}
