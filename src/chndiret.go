package main

import "fmt"

/* 发送 chan<- 作为左值 */
func ping(pings chan<- string, msg string) {
	pings <- msg
}

/* <-chan 作为右值  */
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	/* 发送信息到通道pings */
	ping(pings, "passed message")

	/* 从pings 读取消息，发送到pongs */
	pong(pings, pongs)

	/* 从pongs 读取信息并打印 */
	fmt.Println(<-pongs)

}
