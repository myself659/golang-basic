package main

import (
	"fmt"
	"time"
)

func checksend(ch chan uint, num uint) {
	select {
	case _, done := <-ch:
		if done == true {
			ch <- num
		} else {
			fmt.Println(ch, ":checksend has closed ")
		}
	default:
		ch <- num
	}
}

func checkclose(ch chan uint) {
	select {
	case _, done := <-ch:
		if done == true {
			close(ch)
			fmt.Println(ch, ":checkclose first close in recv")
		} else {
			fmt.Println(ch, ":checkclose has closed in recv")
		}
	default:
		fmt.Println(ch, ":checkclose first close in default")
		close(ch)
	}
}

func main() {
	var size uint = 1
	for ; size < 1000; size++ {
		ch := make(chan uint, size)
		// 可以这样使用 等同 ch := make(chan bool)
		//ch := make(chan bool)
		go func(ch chan uint) {
			for v := range ch {
				v = v
			}
		}(ch)
		go checksend(ch, size)
		go checksend(ch, size)
		go checkclose(ch)
		go checkclose(ch)
	}

	fmt.Println("exit")
	<-time.After(time.Second)
}

/*

0xc820a99000 :checkclose first close in default
0xc820a99000 :checkclose has closed in recv
panic: close of closed channel

goroutine 3383 [running]:
panic(0x4ba060, 0xc82000a4b0)
        /usr/local/go/src/runtime/panic.go:464 +0x3e6
main.checkclose(0xc820a8a800)
        /share/go/src/chancloseclose.go:32 +0x3e9
created by main.main
        /share/go/src/chancloseclose.go:49 +0xed
exit status 2

0xc8208c2a00 :checkclose first close in default
0xc8208c2a00 :checkclose first close in default
为什么会出现这种情况，两个核并发并不安全
*/
