package main

import "fmt"
import (
	"strconv"
)

func main() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1=", 1+1)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	s := "1.234566"
	f, _ := strconv.ParseFloat(s, 2)
	fmt.Println(f)
	sf := fmt.Sprintf("%0.2f", 1.23456)
	f, _ = strconv.ParseFloat(s, 4)
	fmt.Println(f)
	f, _ = strconv.ParseFloat(sf, 4)
	fmt.Println(f)
	d := 1
	fmt.Printf("%06d", d)
	fmt.Printf("%6d", d)
	d = 123456
	fmt.Printf("%06d", d)
	// fmt.Printf("%x6d", d)
}
