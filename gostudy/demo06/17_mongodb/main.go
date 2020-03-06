package main

import (
	"fmt"
)

func main() {
	defer cancel()
	//insertOne()
	//insertMany()
	//selectOne()
	//selectAll()
	//delete()
	update()
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
