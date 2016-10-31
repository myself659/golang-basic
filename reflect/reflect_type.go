package main

import (
	"fmt"
	"reflect"
)

func main() {
	// iterate through the attributes of a Data Model instance
	for name, mtype := range attributes(Dish{}) {
		fmt.Printf("Name: %s, Type: %s\n", name, mtype.Name())
	}

	for name, mtype := range attributes(&Dish{}) {
		fmt.Printf("Name: %s, Type: %s\n", name, mtype.Name())
	}
}

// Data Model
type Dish struct {
	Id     int
	Name   string
	Origin string
	Query  func() //函数是一种类型，就当函数指针吧
}

// Example of how to use Go's reflection
// Print the attributes of a Data Model
func attributes(m interface{}) map[string]reflect.Type {
	typ := reflect.TypeOf(m)
	// if a pointer to a struct is passed, get the type of the dereferenced object
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	// create an attribute data structure as a map of types keyed by a string.
	attrs := make(map[string]reflect.Type)
	// Only structs are supported so return an empty result if the passed object
	// isn't a struct
	if typ.Kind() != reflect.Struct {
		fmt.Printf("%v type can't have attributes inspected\n", typ.Kind())
		return attrs
	}

	// loop through the struct's fields and set the map
	for i := 0; i < typ.NumField(); i++ {
		// 反射是怎么一个结构体  有多少个结构成员，每一个成员分为对应名字，对应的数据类型
		p := typ.Field(i)
		if !p.Anonymous {
			attrs[p.Name] = p.Type
		}
	}

	return attrs
}

/*

让实现类型无关
实现运行时获取类型

*/
