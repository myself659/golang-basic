package benchdemo

import (
	_"fmt"
)

var dmap = make(map[string]int)

var imap = make(map[interface{}]interface{})

func PrintOnly() {
	j := 0
	for i := 0; i < 1000; i++ {
		//fmt.Println("hello")
		j++
	}
}

func PrintByGoRoutine() { //为什么这个
	for i := 0; i < 1000; i++ {
		j := 0
		go func() {
			//fmt.Println("hello")
			j++
		}()
	}
}

func AddDMap(k string, v int) {
	dmap[k] = v
}

func AddIMap(k string, v int) {
	imap[k] = v
}

func Concat(ss ...string) string {
	var r string
	for _, s := range ss {
		r += s
	}
	return r
}
