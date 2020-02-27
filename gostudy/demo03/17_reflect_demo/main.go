package main

import (
	"fmt"
	"reflect"
)

type cat struct{}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
	fmt.Printf("name:%v, kind:%v\n", v.Name(), v.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	switch v.Kind() {
	case reflect.Int64:
		fmt.Printf("type is int64, value:%v\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is float32, value:%v\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64, value:%v\n", float64(v.Float()))
	}
}

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200)
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func main() {
	var a float32 = 3.14
	reflectType(a)
	var b int64 = 100
	reflectType(b)
	var c cat
	reflectType(c)

	reflectValue(a)
	reflectValue(b)
	reflectSetValue1(&b)
	fmt.Println(b)
	reflectSetValue2(&b)
	fmt.Println(b)
}
