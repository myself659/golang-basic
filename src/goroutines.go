package main

import "fmt"

func done(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	done("start")
	go done("goroutine ")

	go func(msg string) {
		fmt.Println(msg)
	}("Anonymous goroutine")

	var input string
	fmt.Scanln(&input)

	fmt.Println("finish")
}

/* Go 运行时是以异步的方式运行协程的  */
