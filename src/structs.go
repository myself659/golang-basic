package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	fmt.Println(person{"bob", 20})

	fmt.Println(person{name: "eric", age: 28})

	s := person{name: "alen", age: 23}
	sp := &s
	fmt.Println(sp.age)
	sp.age = 24
	fmt.Println(sp.age)

}
