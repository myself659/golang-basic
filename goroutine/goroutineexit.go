package main

import (
	"fmt"
	"time"
)

func TenTicks(ticker *time.Ticker) {
	var i int = 0
	for t := range ticker.C {
		i++
		fmt.Println(t)
		if i == 10 {
			fmt.Println("10 ticks done")

			break
		}
	}
}

func main() {
	ticker := time.NewTicker(1 * time.Second)
	go TenTicks(ticker)
	defer ticker.Stop()
	time.Sleep(5 * time.Second)

}
