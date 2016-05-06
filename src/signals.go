package main

import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {

	sigs := make(chan os.Signal, 1) // 用于接收操作系统
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) //注册sigs通道接收SIGINT，SIGTERM信号

	go func() { // 协程完成信号的处理
		sig := <-sigs

		fmt.Println()
		fmt.Println(sig)

		done <- true //通知主进程完成信号处理
	}()

	fmt.Println("awaiting signal")
	<-done //读取完成信号
	fmt.Println("exiting")

}
