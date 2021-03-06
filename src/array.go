package main

import "fmt"

func main() {
	flex := [...]int{1, 2, 4}
	flex = [3]int{1: 11, 2: 22}
	fmt.Println(flex)
	fmt.Println(len(flex))
	var a [5]int
	fmt.Println("empty:", a)
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("getlen:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var two [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			two[i][j] = j
		}
	}

	fmt.Println("two:", two)
}
