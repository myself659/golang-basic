package main

import "errors"
import "fmt"

func f1(arg int) (int, error) {
	if arg == 45 {
		return -1, errors.New("can not work with 45")
	}

	return arg + 5, nil
}

type argError struct {
	arg  int
	prob string
}

/* 使用Error方法自定义error类型 */
func (e *argError) Error() string {
	return fmt.Sprintf("%d -%s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 45 {
		return -1, &argError{arg, "can not work with it"}
	}

	return arg + 3, nil
}

func main() {
	for _, i := range []int{30, 45} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worker:", r)
		}
	}

	for _, i := range []int{30, 45} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worker:", r)
		}

	}
	/* e为argError的指针，对应f2函数的返回值  */
	_, e := f2(45)
	//fmt.Println(e.arg)
	//fmt.Println(e.prob)

	/* ae, ok := e.(*argError); 理解  */
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

}
