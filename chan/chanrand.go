package main

import (
	"fmt"
	"time"
)

func sendAward(ch chan string) {

	ch <- "you get the award "

	close(ch)
}

func recvAward(ch chan string, id int) {
	s, ok := <-ch
	if ok == false {
		fmt.Println(id, ":sorry,you are not get the award")
	} else {
		fmt.Println(id, ":", s)
	}
}
func main() {
	ch := make(chan string)
	go sendAward(ch)
	for i := 0; i < 5; i++ {
		go recvAward(ch, i)
	}

	<-time.After(time.Second)
}

/*
并不能保证随机
*/
