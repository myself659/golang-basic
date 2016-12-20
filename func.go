package main

import (
	"fmt"
)

func loopdone(f func(i int)) {
	for i := 0; i < 3; i++ {
		f(i)
	}
}

func main() {
	var sum int
	loopdone(func(i int) {
		sum += i
		fmt.Println(sum)
	})

}
