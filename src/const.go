package main

import "fmt"
import "math"

const s string = "constant"

type Point struct {
	x int
	y int
}

const (
	MONDAY  = itoa
	TUESDAY = itoa
)

//const start = Point{x: 0, y: 0}

func main() {
	var fn float64 = 1.23456748
	fmt.Printf("%.3f", fn)

	fmt.Println(s)

	const n = 500000

	const d = 60

	fmt.Println(d)

	fmt.Println(math.Sin(n)) //区分大小写
}
