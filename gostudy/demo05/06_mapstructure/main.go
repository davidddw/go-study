package main

import (
	"fmt"

	maps "github.com/mitchellh/mapstructure"
)

type user struct {
	ID      int      `db:"id"`
	Name    string   `db:"name"`
	Age     float64  `db:"age"`
	Sex     int      `db:"sex"`
	Married bool     `db:"married"`
	Hobby   []string `db:"hobby"`
}

type department struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Info string `db:"dinfo"`
}

func main() {
	map1 := &map[string]interface{}{
		"id":    2,
		"name":  "zhiming",
		"dinfo": "kaifa",
	}
	map2 := &map[string]interface{}{
		"id":      2,
		"name":    "zhiming",
		"age":     20.2,
		"sex":     1,
		"married": true,
		"hobby":   [3]string{"eat", "eat1", "eat1"},
	}
	var u user
	//u := new(user)
	convertMapToUser(map2, &u)
	fmt.Printf("%#v\n", u)
	var u2 department
	//u2 := new(department)
	convertMapToUser(map1, &u2)
	fmt.Printf("%#v\n", u2)
}

func convertMapToUser(data *map[string]interface{}, out interface{}) {
	cfg := &maps.DecoderConfig{
		Metadata: nil,
		Result:   out,
		TagName:  "db",
	}
	decoder, _ := maps.NewDecoder(cfg)
	decoder.Decode(*data)
}
