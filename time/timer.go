package main

import "time"
import "fmt"

func main() {
	/* 一次性定时器 */
	timer1 := time.NewTimer(time.Second * 2)

	/* 直到这个定时器的通道 C 明确的发送了定时器失效的值之前，将一直阻塞 */
	<-timer1.C
	fmt.Println("timer1 expired")

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("timer2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer2 stopped")
	}

}
