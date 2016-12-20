package main

import (
	"fmt"
)

type Handle interface {
	Server(s string)
}

func server1(s string) {
	fmt.Println("server1:", s)
}

func server2(s string) {
	fmt.Println("server2:", s)
}

func main() {

}
