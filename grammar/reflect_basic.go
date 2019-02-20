package grammar

import (
	"fmt"
	"reflect"
)

type product struct {
	id    int
	name  string
	price int
	stock *map[string]int
}

func createQuery(i interface{}) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	fmt.Println("type:", t, ", type.Kind:", t.Kind(), ", type.Name:", t.Name())
	fmt.Println("value:", v, ", value.Kind:", v.Kind(), ", value.Num", v.NumField())

	for i := 0; i < v.NumField(); i++ {
		fmt.Println("field:", i, ", name:", t.Field(i).Name, ", type:", t.Field(i).Type)
		switch v.Field(i).Kind() {
		case reflect.Int:
			fmt.Printf("v:%d, can set:%t\n", v.Field(i).Int(), v.Field(i).CanSet())
			//v.Field(i).SetInt(123)
		case reflect.String:
			fmt.Printf("v:%s, can set:%t\n", v.Field(i).String(), v.Field(i).CanSet())
		case reflect.Ptr:
			mapVal := v.Field(i).Elem()
			fmt.Println("v:", mapVal, ", can set:", v.Field(i).CanSet())
			fmt.Println(mapVal.MapKeys())
		default:
			fmt.Println("Unsupported type")
			return
		}
	}
}

func BasicReflect() {
	o := product{
		id:   456,
		name: "foobar",
		stock: &map[string]int{
			"foo": 1,
			"bar": 2,
		},
	}
	createQuery(o)
}
