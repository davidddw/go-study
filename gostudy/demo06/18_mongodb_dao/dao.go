package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient  *mongo.Client
	mongoContext context.Context
)

const (
	// DbName table name
	DbName = "mydb"
	// Username user
	Username = "cloud"
	// Password passwd
	Password = "passwd"
	// Host host ip
	Host = "localhost"
)

func init() {
	mongoContext, _ = context.WithTimeout(context.Background(), 30*time.Second)
	opts := &options.ClientOptions{}
	opts.SetAuth(options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		AuthSource:    DbName,
		Username:      Username,
		Password:      Password,
	}).ApplyURI(fmt.Sprintf("mongodb://%s:27017", Host))
	client, err := mongo.Connect(mongoContext, opts)
	if err != nil {
		fmt.Println(err)

	}
	mongoClient = client
}

// GetSession return client and context
func GetSession() (*mongo.Client, context.Context) {
	return mongoClient, mongoContext
}

func getCollect(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database(DbName).Collection(collectionName)
}
