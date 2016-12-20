package main

import "os"
import "strings"
import "fmt"

const oct = 077
const hex = 0x99
const ff = 1e2
const num = ^0

func main() {
	//fmt.Printf("%d\n", uint8(2.5))
	fmt.Printf("%d\n", uint8(35.0))
	fmt.Printf("%d\n", num)
	fmt.Printf("%f\n", ff)
	fmt.Printf("%d\n", oct)
	fmt.Printf("%d\n", hex)
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
