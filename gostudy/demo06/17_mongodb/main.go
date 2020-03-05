package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Book entity
type Book struct {
	ID       string
	Name     string
	Category string
	Weight   int
	Author   AuthorInfo
}

// AuthorInfo author
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
			ID:       "1",
			Name:     "深入理解计算机操作系统",
			Category: categoryComputer,
			Weight:   1,
			Author: AuthorInfo{
				Name:    "兰德尔 E.布莱恩特",
				Country: countryAmerica,
			},
		},
		&Book{
			ID:       "2",
			Name:     "深入理解Linux内核",
			Category: categoryComputer,
			Weight:   1,
			Author: AuthorInfo{
				Name:    "博韦，西斯特",
				Country: countryAmerica,
			},
		},
		&Book{
			ID:       "3",
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

// GetID inpletment
func (e Book) GetID() string {
	return e.ID
}

// SetID setter
func (e *Book) SetID(id string) {
	e.ID = id
}

func main() {
	defer cancel()
	// book1 := &Book{
	// 	ID:       "4",
	// 	Name:     "三体",
	// 	Category: categorySciFi,
	// 	Weight:   1,
	// 	Author: AuthorInfo{
	// 		Name:    "刘慈欣",
	// 		Country: countryChina,
	// 	},
	// }
	//fmt.Println(Mongo.Insert("book", book1))
	var one Book
	Mongo.GetOneByID("book", "18a0a36a-8048-4a14-b555-1ae0a70e5874", &one)
	fmt.Println(one)

	var all []*Book
	skip, limit := int64(0), int64(5)
	filter := PageFilter{
		SortBy:   "name",
		SortMode: 1,
		Skip:     &skip,
		Limit:    &limit,
	}
	Mongo.GetAllByFilter("book", filter, &all)
	for _, v := range all {
		fmt.Println(v)
	}

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
