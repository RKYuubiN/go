package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func main() {
	fmt.Println("Reflection")

	foo := Foo{
		Name:    "Shraddha",
		Address: "Somewhere in the middle",
	}

	data, err := structToMap(foo, "json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)

	// fmt.Println(myfoo)
	// data[myf]
	// structToMap(foo)
}

func structToMap(data interface{}, tag string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(data)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("only accept struct; got %T", v)
	}

	typ := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fi := typ.Field(i)
		if tagv := fi.Tag.Get(tag); tagv != "" {
			out[tagv] = v.Field(i).Interface()
		}
	}

	return out, nil
}
