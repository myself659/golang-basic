package main

import "fmt"

/* 多个函数返回值  */
func multiret() (int, int) {
	return 3, 7
}

func main() {
	/* 获取多个返回值  */
	a, b := multiret()

	fmt.Println(a)
	fmt.Println(b)

	/* -, 表示默认变量 */
	_, c := multiret()
	fmt.Println(c)

}
