package grcost

import (
	"fmt"
)

func PrintOnly() {
	for i := 0; i < 1000; i++ {
		fmt.Println("hello")
	}
}

func PrintByGoRoutine() {
	for i := 0; i < 1000; i++ {
		go func() {
			fmt.Println("hello")
		}()
	}
}
