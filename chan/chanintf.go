package main

import (
	"fmt"
	_ "io"
	"reflect"
)

type Point struct {
	a int
	b int
}

func (self *Point) String() string {

	return fmt.Sprintf("a=%d, b=%d", self.a, self.b)
}
func main() {
	ch := make(chan interface{})
	go func() {
		ch <- 1
		ch <- "abc"
		temp := Point{a: 1, b: 2}
		fmt.Printf("%s\n", temp.String())
		ch <- temp
		close(ch)
	}()

	for intf := range ch {
		rt := reflect.TypeOf(intf)
		fmt.Println("reflect.TypeOf:", rt)
		fmt.Println("reflect type kind:", rt.Kind())
		rv := reflect.ValueOf(intf)
		fmt.Println("reflect.ValueOf:", rv)
		//fmt.Println("reflect.ValueOf value:", rv.)
		fmt.Println(intf)
		//fmt.Printf("%s\n", intf.String())
	}

}

/*

如何实现用户自定义chan 类型

*/
