package main

import (
	"fmt"
)

func main() {
	ch := make(chan interface{})
	go func() {
		ch <- "ab"
	}()

	var temp int
	temp <- ch // 编译阶段就能检查出来
	fmt.Println(temp)
}
