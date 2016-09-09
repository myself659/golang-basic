package main

import (
	"fmt"
	"runtime"
)

func test(s string, x int) (r string) {
	r = fmt.Sprintf("test: %s, %d", s, x)
	runtime.Breakpoint()

	return r
}

func main() {
	s := "haha"
	i := 100
	fmt.Println(test(s, i))
}
