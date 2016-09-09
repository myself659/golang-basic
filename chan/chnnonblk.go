package main

import "fmt"

func main() {

	message := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-message:
		fmt.Println("received msg", msg)
	default: /* 实现非阻塞操作 */
		fmt.Println("no msg received")
	}

	msg := "hi"
	select {
	case message <- msg:
		fmt.Println("sent msg", msg)
	default:
		fmt.Println("no msg sent")
	}

	select {
	case msg := <-message:
		fmt.Println("received msg", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
