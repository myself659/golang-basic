package main

import (
	"fmt"
	"time"
)

func Notify() {
	var ch chan int
	for i := 0; i < 2; i++ {
		ch = make(chan int)
		for j := 0; j < 3; j++ {
			go Recv(ch)
		}
		<-time.After(1 * time.Second)
		close(ch)

	}
}

func Recv(ch chan int) {
	var rch chan int
	//fmt.Println(ch)
	rch = ch
	for {
		select {
		case _, ok := <-rch:
			if ok == false {
				fmt.Println(rch, "close ")
				//go Recv(ch)
				return
			}

		}
	}
}

func main() {
	go Notify()

	<-time.After(5 * time.Second)
}
