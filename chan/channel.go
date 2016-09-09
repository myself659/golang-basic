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

	count := make(chan int)
	/* goroutine 随机运行 */
	go func() { count <- 1 }()
	go func() { count <- 2 }()
	go func() { count <- 3 }()
	go func() { count <- 4 }()

	for {
		temp, ok := <-count
		fmt.Println(temp)
		if !ok {
			fmt.Println("chan drain")
			return
		}
	}

	for i := 0; i < 4; i++ { /* 次数不致也会导致死锁 如何避免这个问题呢 */
		temp := <-count
		fmt.Println(temp)
	}

	/*
		for {
			temp := <-count
			fmt.Println(temp)
		}
	*/

}
