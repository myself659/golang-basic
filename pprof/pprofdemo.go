package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	chint := make(chan int)
	tick := time.NewTicker(time.Second)
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	go func() {
		for i := 0; i < 4000000; i++ {
			chint <- i
		}
	}()

	cnt := 0
	for {
		select {
		case value := <-chint:
			//fmt.Println(value)
			cnt++
			_ = value
		case tt := <-tick.C:
			fmt.Println(tt, cnt)
		}
	}
}
