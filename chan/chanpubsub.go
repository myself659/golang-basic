package main

import (
	"fmt"
	"strconv"
	"time"
)

type conttype string
type suber chan<- conttype

var addsuber = make(chan suber)
var delsuber = make(chan suber)
var contents = make(chan conttype)
var stoppub = make(chan bool)

//只完成数据交付  数据交付方式 slice 不断从中获取
func procsubrsp(id int, ch <-chan conttype) {
	for content := range ch {
		fmt.Println("id:", id, content)
	}
	fmt.Printf("id %d: suber exit\n", id)
}

func addsubsciber(id int) suber {
	ch := make(chan conttype)

	addsuber <- ch
	go procsubrsp(id, ch)
	return ch
}

func delsubsciber(sub suber) {
	delsuber <- sub
}

// 作为公用函数，作为包的函数
func publisher() {

	subers := make(map[suber]bool)
	for {
		select {
		case content := <-contents:
			//发布订阅信息给订阅者
			for sub := range subers {
				sub <- content
			}
		case sub := <-addsuber:
			subers[sub] = true
		case sub := <-delsuber:
			delete(subers, sub)
			close(sub)
		case <-stoppub:
			for sub := range subers {
				delete(subers, sub)
				close(sub)
			}
			fmt.Println("stop pub\n")
		}
	}

}

func main() {

	ticker := time.NewTicker(time.Second)
	go func() {
		var prefix conttype = "pub_"
		i := 0
		for range ticker.C {
			i++

			contents <- prefix + conttype(strconv.Itoa(i))
		}
	}()
	go publisher()
	i := 0
	ch1 := addsubsciber(i)
	i++
	addsubsciber(i)
	i++
	addsubsciber(i)

	<-time.After(3 * time.Second)
	delsubsciber(ch1)

	<-time.After(8 * time.Second)
	ticker.Stop()
	stoppub <- true

	<-time.After(3 * time.Second)

}
