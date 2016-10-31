package main

import (
	"fmt"
	"reflect"
)

type Point struct {
	x    int
	y    int
	name string
}

func printint(i int) {
	fmt.Println("printi:", i)
}

func PrintPoint(i interface{}) {
	rtype := reflect.TypeOf(i)
	fmt.Println(rtype)
	rk := rtype.Kind()
	rv := reflect.ValueOf(i) // 从struct对应reflect.Value取值field
	fmt.Println(rk)
	//fmt.Println(rv)
	fmt.Println(rv.NumField())
	for i := 0; i < rv.NumField(); i++ {
		rfield := rv.Field(i)
		fmt.Println(rfield)
		// 获得filed真实类型
		rft := rfield.Type()
		fmt.Println(rft)

		switch rft.Kind() { // 反射值对应的类型
		case reflect.Int:
			{
				//调用接口转化为string
				var iv int
				// iv = int(reflect.ValueOf(rfield).Int()) 错误用法
				iv = int(rfield.Int())
				printint(iv)

			}
		case reflect.String:
			{
				fmt.Printf("str:%s\n", rfield.String())
			}

		}

		rrf := reflect.TypeOf(rfield)

		fmt.Println(rrf.Kind())
		rrv := reflect.ValueOf(rfield)
		//printint(rrv)
		fmt.Println(rrv.Type())
		fmt.Println(rfield)
	}

}

func main() {
	p := Point{x: 1, y: 10, name: "abc"}
	PrintPoint(p)
}
