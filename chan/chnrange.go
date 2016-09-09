package main

import "fmt"

func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"

	close(queue)

	/* 关闭通道后仍能读取消息 */
	for elem := range queue {
		fmt.Println(elem)
	}
}
