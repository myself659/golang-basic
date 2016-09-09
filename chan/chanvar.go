package main

import (
	"fmt"
	"time"
)

type Wrapper struct {
	ch chan int
}

func Notify(w *Wrapper) {

	for i := 0; i < 2; i++ {
		w.ch = make(chan int)
		fmt.Println(w.ch, "create")
		for j := 0; j < 3; j++ {
			go Recv(w)
		}
		<-time.After(1 * time.Second)
		close(w.ch)

	}
}

func Recv(w *Wrapper) {
	var rch chan int
	rch = nil // explicit
	for {
		if rch != w.ch {
			rch = w.ch
		} else {
			rch = nil
		}
		fmt.Println("new:", rch)
		select {
		case _, ok := <-rch:
			if ok == false {
				fmt.Println(rch, ":close ")
				//go Recv(ch)

				continue // 跳出select，重新select
				// break 与continue 的差别
			}

		}
	}
}

func main() {
	w := new(Wrapper)
	go Notify(w)

	<-time.After(5 * time.Second)
}
