package main

import "time"
import "fmt"

func main() {
	c1 := make(chan string, 1)

	/* 执行协程 */
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1: /* 读取到c1通道内容  */
		fmt.Println(res)
	case <-time.After(time.Second * 1): /* 等待1s */
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3): /* 等待3s */
		fmt.Println("timeout 2")
	}

}
