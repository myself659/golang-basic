package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var ops int64

	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	/* 读取通道信息 回应指定通道 */
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
				//write.resp = true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				/* 结构体初始化与取地址 */
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)} // }不可另占一行
				reads <- read
				<-read.resp
				fmt.Println("in read", time.Now())
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				fmt.Println("in write", time.Now())
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	time.Sleep(time.Second) /* 让协程运行1s */

	opsFinal := atomic.LoadInt64(&ops)

	fmt.Println("ops;", opsFinal)

}
