package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Book struct {
	Id       primitive.ObjectID `bson:"_id"`
	Name     string
	Category string
	Weight   int
	Author   AuthorInfo
}

type AuthorInfo struct {
	Name    string
	Country string
}

const (
	categoryComputer = "计算机"
	categorySciFi    = "科幻"
	countryChina     = "中国"
	countryAmerica   = "美国"
)

var (
	books = []interface{}{
		&Book{
			Id:       primitive.NewObjectID(),
			Name:     "深入理解计算机操作系统",
			Category: categoryComputer,
			Weight:   1,
			Author: AuthorInfo{
				Name:    "兰德尔 E.布莱恩特",
				Country: countryAmerica,
			},
		},
		&Book{
			Id:       primitive.NewObjectID(),
			Name:     "深入理解Linux内核",
			Category: categoryComputer,
			Weight:   1,
			Author: AuthorInfo{
				Name:    "博韦，西斯特",
				Country: countryAmerica,
			},
		},
		&Book{
			Id:       primitive.NewObjectID(),
			Name:     "三体",
			Category: categorySciFi,
			Weight:   1,
			Author: AuthorInfo{
				Name:    "刘慈欣",
				Country: countryChina,
			},
		},
	}
)

func main() {
	defer cancel()
	// book1 := &Book{
	// 	Id:       primitive.NewObjectID(),
	// 	Name:     "三体",
	// 	Category: categorySciFi,
	// 	Weight:   1,
	// 	Author: AuthorInfo{
	// 		Name:    "刘慈欣",
	// 		Country: countryChina,
	// 	},
	// }
	//fmt.Println(Mongo.Create("book", book1))
	skip, limit := int64(0), int64(5)
	filter := PageFilter{
		SortBy:   "name",
		SortMode: 1,
		Skip:     &skip,
		Limit:    &limit,
	}
	fmt.Println(Mongo.Get("book", "5e60c5e30aa3d18cce90a512"))
	fmt.Println(Mongo.List("book", filter))
	// skip, limit := int64(0), int64(5)
	// filter := make(map[string]interface{})
	// filter["weight"] = "1"
	//fmt.Println(Mongo.Count("book", PageFilter{}))
	//getData()
}

func getData() {
	opts := &options.ClientOptions{}
	opts.SetAuth(options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		AuthSource:    db,
		Username:      user,
		Password:      pass,
	}).ApplyURI(fmt.Sprintf("mongodb://%s:27017", host))
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return
	}
	collection := client.Database("mydb").Collection("book")
	var one Book
	objID, _ := primitive.ObjectIDFromHex("5e60c5e30aa3d18cce90a512")
	fmt.Println(collection.FindOne(context.Background(), bson.M{"_id": objID}))
	err = collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&one)
	if err != nil {
		return
	}
	fmt.Println("collection.FindOne: ", one)
}
