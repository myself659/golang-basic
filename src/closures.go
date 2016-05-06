package main

import "fmt"

//
func intseq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	// 自动推导类型
	nextInt := intseq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := intseq()

	fmt.Println(newInt())

}
