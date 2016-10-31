package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 4.1

	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value:", v.Float())

	type t struct {
		N int
	}
	var n = t{42}
	// N at start
	fmt.Println(n.N)
	// pointer to struct - addressable
	pn := reflect.ValueOf(n)
	fmt.Println("reflect.ValueOf(n):", pn)
	ps := reflect.ValueOf(&n)
	fmt.Println("reflect.ValueOf(&n):", ps)
	_ = ps
	// struct
	// 需要修改结构体的成员值，需要用指针进行反射
	s := ps.Elem()
	fmt.Println("Elem:", s)
	// s := pn.Elem()
	if s.Kind() == reflect.Struct {
		// exported field
		f := s.FieldByName("N")
		if f.IsValid() {
			// A Value can be changed only if it is
			// addressable and was not obtained by
			// the use of unexported struct fields.
			if f.CanSet() {
				// change value of N
				if f.Kind() == reflect.Int {
					x := int64(7)
					if !f.OverflowInt(x) {
						f.SetInt(x)
					}
				}
			}
		}
	}
	// N at end
	fmt.Println(n.N)
}
