package main

import (
	"fmt"
	"time"
)

func create() {
	ch := make(chan int)
	close(ch)
}

func main() {
	for i := 0; i < 1000000; i++ {
		go create()
		if i%100 == 0 {
			time.Sleep(time.Second)
		}
	}

	<-time.After(100 * time.Second)
}
