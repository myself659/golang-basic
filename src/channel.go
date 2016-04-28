package main

import "fmt"

func main() {
	//
	message := make(chan string)
	foomsg := make(chan string)

	go func() { message <- "ping" }() // write string to message chan
	go func() { foomsg <- "foo" }()   // write  string to foomsg chan

	msg := <-message // read chan to var  msg

	fmt.Println(msg)

	foo := <-foomsg

	fmt.Println(foo)
}
