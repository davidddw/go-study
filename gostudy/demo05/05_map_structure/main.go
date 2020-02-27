package main

import (
	"reflect"
	"strconv"
)

type user struct {
	ID      int      `db:"id"`
	Name    string   `db:"name"`
	Age     float64  `db:"age1"`
	Sex     int      `db:"sez"`
	Married bool     `db:"married"`
	Hobby   []string `db:"hobby"`
}

type teacher struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func convertMapToObject(record *map[string]string) user {
	var u user
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(&u).Elem() // 为了改变对象的内部值，需使用引用
	for i := 0; i < t.NumField(); i++ {
		f := v.Field(i)
		fieldName := t.Field(i).Tag.Get("db")
		if f.Kind() == reflect.Int {
			val, _ := strconv.Atoi((*record)[fieldName]) // 通过tag获取列数据
			f.SetInt(int64(val))
		} else if f.Kind() == reflect.String {
			f.SetString((*record)[fieldName])
		}
	}
	return u
}

// MapToStruct convert map[string]string to structure type by tag
func MapToStruct(data *map[string]string, out interface{}) {
	v := reflect.ValueOf(out).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		val := (*data)[t.Field(i).Tag.Get("db")]
		name := t.Field(i).Name
		switch v.Field(i).Kind() {
		case reflect.String:
			v.FieldByName(name).SetString(val)
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
			i, err := strconv.Atoi(val)
			if err != nil {
				continue
			}
			v.FieldByName(name).SetInt(int64(i))
		case reflect.Uint16, reflect.Uint32, reflect.Uint64:
			i, err := strconv.Atoi(val)
			if err != nil {
				continue
			}
			v.FieldByName(name).SetUint(uint64(i))
		case reflect.Float32, reflect.Float64:
			f, err := strconv.ParseFloat(val, 64)
			if err != nil {
				continue
			}
			v.FieldByName(name).SetFloat(f)
		default:
			//logger.ERROR(LOG_NAME, "unknown type:%+v", ss.Field(i).Kind())
		}
	}
	return
}

// DataToStruct convert map[string]interface{} to structure type by tag
func DataToStruct(data *map[string]interface{}, out interface{}) {
	v := reflect.ValueOf(out).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		val := (*data)[t.Field(i).Tag.Get("db")]
		name := t.Field(i).Name
		switch v.Field(i).Kind() {
		case reflect.String:
			v.FieldByName(name).SetString(val.(string))
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
			v.FieldByName(name).SetInt(int64(val.(int)))
		case reflect.Uint16, reflect.Uint32, reflect.Uint64:
			v.FieldByName(name).SetUint(val.(uint64))
		case reflect.Float32, reflect.Float64:
			v.FieldByName(name).SetFloat(val.(float64))
		case reflect.Bool:
			v.FieldByName(name).SetBool(val.(bool))
		default:
			//logger.ERROR(LOG_NAME, "unknown type:%+v", ss.Field(i).Kind())
		}
	}
	return
}
