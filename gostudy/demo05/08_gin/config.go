package main

import (
	"errors"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// OracleConfig mysql server config
type OracleConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	Sid      string `ini:"sid"`
}

// Config ...
type Config struct {
	OracleConfig `ini:"oracle"`
}

func checkIni(filename string, data interface{}) (lines []string, err error) {
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err = errors.New("data should be a pointer")
		return
	}
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data should be a struct")
		return
	}
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	lines = strings.Split(string(buf), "\r\n")
	return
}

// LoadIni load ini file
func LoadIni(filename string, data interface{}) (err error) {
	t := reflect.ValueOf(data).Type().Elem()
	v := reflect.ValueOf(data).Elem()
	lines, err := checkIni(filename, data)
	var structName string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") || len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = errors.New("syntax error")
				return
			}
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = errors.New("syntax error")
				return
			}
			for i := 0; i < t.NumField(); i++ {
				field := t.Field(i)
				if sectionName == field.Tag.Get("ini") {
					structName = field.Name
					// fmt.Println("sec:", structName)
				}
			}
		} else {
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = errors.New("syntax error")
				return
			}
			index := strings.Index(line, "=")
			iniKey := strings.TrimSpace(line[:index])
			iniValue := strings.TrimSpace(line[index+1:])

			if t.Kind() != reflect.Struct {
				err = errors.New("data should be a struct")
				return
			}
			var fieldType reflect.StructField
			sValue := v.FieldByName(structName)
			for i := 0; i < sValue.NumField(); i++ {
				field := sValue.Type().Field(i)
				if field.Tag.Get("ini") == iniKey {
					fieldType = field
					break
				}
			}
			fieldObj := sValue.FieldByName(fieldType.Name)
			switch fieldType.Type.Kind() {
			case reflect.String:
				fieldObj.SetString(iniValue)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(iniValue, 10, 64)
				if err != nil {
					return
				}
				fieldObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(iniValue)
				if err != nil {
					return
				}
				fieldObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				var valueFloat float64
				valueFloat, err = strconv.ParseFloat(iniValue, 64)
				if err != nil {
					return
				}
				fieldObj.SetFloat(valueFloat)
			}
		}
	}
	return
}
