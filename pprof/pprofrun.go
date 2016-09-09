package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var cnt int = 0

func fib(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	return fib(n-1) + fib(n-2)
}

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	go fib(10000)
	go fib(10000)
	go fib(100000)
	go fib(10000000)

	ticker := time.NewTicker(time.Second)
	for tt := range ticker.C {
		cnt++
		fmt.Println(tt, cnt)
		if cnt == 200 {
			ticker.Stop()
			break
		}
	}
	fmt.Println("stop")
}
