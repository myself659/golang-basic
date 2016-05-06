package main

import "time"
import "fmt"

func main() {

	reqs := make(chan int, 5)
	for i := 1; i < 5; i++ {
		reqs <- i
	}

	close(reqs)

	limiter := time.Tick(time.Millisecond * 200)

	for tempreq := range reqs {
		<-limiter /* 等待200ms 超时 */
		fmt.Println("req", tempreq, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now() /* 立即写入三个消息 */
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t /* 每200ms 向限速器写一个消息 */
		}
	}()

	burstyReqs := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		burstyReqs <- i
	}

	close(burstyReqs)
	for tempreq := range burstyReqs {
		<-burstyLimiter
		fmt.Println("burstyReq", tempreq, time.Now())
	}

}

/* 通过通道长度 和定时器控制写入消息速度 来实现限速 */
