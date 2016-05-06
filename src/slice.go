package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	//s[3] = "d"  //index out of range  拒绝访问溢出问题

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println("appd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		len := i + 1
		twoD[i] = make([]int, len)
		for j := 0; j < len; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twod:", twoD)

}
