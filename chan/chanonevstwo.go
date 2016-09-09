package main

import (
	"fmt"
	"time"
)

var chsyn = make(chan int)
var chbuf = make(chan int, 3)

func send(i int) {
	chsyn <- i
	chbuf <- i
	time.Sleep(1 * time.Second)
}
func recv1() {
	for {
		select {
		case temp := <-chsyn:
			fmt.Println("recv1 chsyn ", temp)
		case temp := <-chbuf:
			fmt.Println("recv1 chbuf ", temp)

		}
	}
}

func recv2() {
	for {

		select {
		case temp := <-chsyn:
			fmt.Println("recv2 chsyn ", temp)
		case temp := <-chbuf:
			fmt.Println("recv2 chbuf ", temp)

		}
	}
}

func main() {
	for i := 0; i < 30; i++ {
		go send(i)
	}

	go recv1()

	go recv1()

	go recv2()

	go recv2()

	time.Sleep(10 * time.Second)
	close(chbuf)
	close(chsyn)
}
