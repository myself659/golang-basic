package main

import "fmt"

/*
先关键字 func
add 函数名
(a int, b int) 函数
int 函数返回值

*/

func add(a int, b int) int {

	return a + b
}

/* 不需要;结束一个语句 */
func main() {
	var c int
	var a int = 1
	var b int = 2
	d := 9

	c = add(a, b)

	fmt.Println("add=", c)
	fmt.Println("d=", d)

}
