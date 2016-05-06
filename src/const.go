package main

import "fmt"
import "math"

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000

	const d = 60

	fmt.Println(d)

	fmt.Println(math.Sin(n)) //区分大小写
}
