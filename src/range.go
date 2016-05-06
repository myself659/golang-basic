package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	sum := 0
	// 空值定义符 _
	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"} // why not map [string][string]
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	// why 将一个字符串转化为map g对应unicode编码
	for i, c := range "golang" {
		fmt.Println(i, c)
	}
}
