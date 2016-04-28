package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"
	//messages <- "test"
	// 消息可以不填充满，但是不超过规格

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	/* 获得消息次数超过发送的消息次数会出现死锁  fatal error: all goroutines are asleep - deadlock!  */
	//fmt.Println(<-messages)

}
