package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		var v int // һ��������ô�ͻ���һ����ʼ����ÿ�����Ͷ�����һ����ʼ��ֵ
		fmt.Println(v)
		v = 5
	}
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j < 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break // must has  condition to go out loop (return )
	}

	for k, j := 0, 10; k < j; k, j = k+1, j-1 {
		fmt.Println(k, j)
	}
}
