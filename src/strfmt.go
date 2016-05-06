package main

import "fmt"
import "os"

type point struct {
	x int
	y int
}

func main() {

	p := point{1, 2}

	a := 2

	// 打印结构体内容
	fmt.Printf("%v\n", p)
	// 打印结构体成员及内容
	fmt.Printf("%+v\n", p)
	// 打印结构体成员及内容及位置
	fmt.Printf("%#v\n", p)
	// 打印结构体类型
	fmt.Printf("%T\n", p)
	// 打印数据类型
	fmt.Printf("%t\n", true)

	fmt.Printf("%t\n", a)

	fmt.Printf("%d\n", 123)

	fmt.Printf("%b\n", 14)

	fmt.Printf("%c\n", 33)

	fmt.Printf("%x\n", 456)

	fmt.Printf("%f\n", 48.9)

	fmt.Printf("%e\n", 12340000000.0)
	fmt.Printf("%E\n", 12340000000.0)

	fmt.Printf("%s\n", "\"string\"")
	fmt.Printf("%q\n", "\"string\"")
	fmt.Printf("%x\n", "hex this")

	fmt.Printf("%p", &p)

	fmt.Printf("|%6d|%6d|\n", 12, 345)

	fmt.Printf("|%6d|%6d|%s|\n", 12, 345, "docker")

	fmt.Printf("|%6.2f|%6.3f|\n", 1.2002, 3.5)

	fmt.Printf("|%6s|%6s|\n", "hello", "mm")

	fmt.Printf("|%-6s|%-6s|\n", "hello", "mm")

	s := fmt.Sprintf("a %s", "string")

	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")

}
