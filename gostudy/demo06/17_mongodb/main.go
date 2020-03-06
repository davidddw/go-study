package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	defer cancel()
	//insertOne()
	//insertMany()
	//selectOne()
	//selectAll()
	//delete()
	//update()
	custom()
}

func custom() {
	var ctx context.Context
	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	opts := &options.ClientOptions{}
	opts.SetAuth(options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		AuthSource:    db,
		Username:      user,
		Password:      pass,
	}).ApplyURI(fmt.Sprintf("mongodb://%s:27017", host))
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		fmt.Println(err)
	}
	collections := client.Database(db).Collection("demo")
	result, err := collections.InsertOne(
		ctx,
		bson.D{
			{Key: "item", Value: "canvas"},
			{Key: "qty", Value: 100},
			{Key: "tags", Value: bson.A{"cotton"}},
			{Key: "size", Value: bson.D{
				{Key: "h", Value: 28},
				{Key: "w", Value: 35.5},
				{Key: "uom", Value: "cm"},
			}},
		})
	if err != nil {
		return
	}
	result, err = collections.InsertOne(
		ctx,
		bson.M{
			"item": "canvas",
			"qty":  100,
			"tags": bson.A{"cotton"},
			"size": bson.M{
				"h":   28,
				"w":   35.5,
				"uom": "cm",
			},
		})
	if err != nil {
		return
	}
	s := result.InsertedID.(primitive.ObjectID).Hex()
	fmt.Println(s)
}

func insertOne() {
	fmt.Println(Mongo.Create("book", book))
}

func insertMany() {
	fmt.Println(Mongo.CreateMany("book", books))
}

func selectOne() {
	var one Book
	Mongo.GetOneByID("book", "5e61d62e39b672ddd6459318", &one)
	fmt.Println(one)
	var two Book
	Mongo.GetOneByUUID("book", "18621960-7fa5-4f31-a917-13e010cc96cc", &two)
	fmt.Println(two)
}

func selectAll() {
	skip, limit := int64(0), int64(5)
	filter := PageFilter{
		SortBy:   "name",
		SortMode: 1,
		Skip:     &skip,
		Limit:    &limit,
	}
	var all []Book
	Mongo.ListByFilter("book", filter, &all)
	for _, v := range all {
		fmt.Println(v)
	}
}

func delete() {
	fmt.Println(Mongo.Delete("book", "5e61d62e39b672ddd6459318"))
}

func update() {
	var book1 Book
	Mongo.GetOneByID("book", "5e61d659aa5b298647fcfd4c", &book1)
	book1.Author.Name = "艾伦A.A.多诺万"
	fmt.Println(Mongo.ModifyByID("book", &book1))
}
