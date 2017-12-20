package main

import "time"
import "fmt"

func main() {

	/* 创建一个500ms的ticker，相当于循环定时器 */
	ticker := time.NewTicker(time.Millisecond * 500)
	/* 每个tick 响应处理 */
	go func() {
		for t := range ticker.C {
			fmt.Println("tick at", t)
		}
	}()
	
	go func() {
		for range ticker.C {
			fmt.Println("tick at")
		}
	}()

	/* 等待睡眠时间 */
	time.Sleep(time.Millisecond * 16000)

	/* 关闭tick */

	ticker.Stop()

	fmt.Println("ticker stopped")
}
