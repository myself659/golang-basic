package main

import "fmt"

var i int
var j = 365.245
var k int = 0
var l, m uint64 = 1, 2
var nanoseconds int64 = 1e9
var inter, floater, stringer = 1, 2.0, "hi"

var (
	vi                          int
	vj                                 = 356.245
	vk                          int    = 0
	vl, vm                      uint64 = 1, 2
	vnanoseconds                int64  = 1e9
	vinter, vfloater, vstringer        = 1, 2.0, "hi"
)

func main() {

	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var e int
	fmt.Println(e)
	// 只能在函数里面
	f := "short"
	fmt.Println(f)
}
