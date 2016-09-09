package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func accessnil() {
	var p *int
	fmt.Println("p =", p)
	*p = 2
}

func main() {

	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGQUIT)
		for {
			<-sigs
			buf := make([]byte, 1<<20)
			runtime.Stack(buf, true)
			log.Printf("===goroutine stack trace ...\n%s\n", buf)
		}
	}()

	//
	for i := 1; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("fmt:", i)
	}
	var temp int = 10
	b := temp
	fmt.Println(b)

	accesserror()

}
