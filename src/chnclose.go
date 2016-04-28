package main

import "fmt"

func main() {

	jobs := make(chan int, 5)

	done := make(chan bool)

	go func() {
		for {
			/* 从jobs接收数据 ，如果接收完成more置为0，表示结束完成 */
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	/* 发送3任务到协程 */
	for j := 1; j <= 5; j++ {
		jobs <- j
		fmt.Println("send job", j)
	}

	close(jobs)

	fmt.Println("send all jobs ")

	//等待接收done通道信息
	<-done
}
