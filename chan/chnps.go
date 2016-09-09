package main

import (
	"fmt"
	"strconv"
	"time"
)

type Content interface{}

type Subscriber chan Content
type Publisher chan<- Content

type ChanBroker struct {
	RegSub      chan Subscriber     // 订阅注册通道
	UnRegSub    chan Subscriber     // 订阅去注册通道
	Contents    chan Content        // 发布通道
	Stop        chan bool           // 停止发布
	Subscribers map[Subscriber]bool // 订阅者MAP
}

//  创建ChanBroker
func NewChanBroker() *ChanBroker {
	ChanBroker := new(ChanBroker)
	ChanBroker.RegSub = make(chan Subscriber)
	ChanBroker.UnRegSub = make(chan Subscriber)
	ChanBroker.Contents = make(chan Content)
	ChanBroker.Stop = make(chan bool)
	ChanBroker.Subscribers = make(map[Subscriber]bool)
	ChanBroker.run()

	return ChanBroker
}

func (self *ChanBroker) run() {

	go func() {
		for {
			select {
			case content := <-self.Contents:
				//发布订阅信息给订阅者
				for sub := range self.Subscribers {
					sub <- content
					// 如果发送Subscriber通道，这时候会报panic，如果是Subscriber没有关闭，只是自己退出可能会导致阻塞
				}
			case sub := <-self.RegSub:
				self.Subscribers[sub] = true
			case sub := <-self.UnRegSub:
				delete(self.Subscribers, sub)
				close(sub)
			case <-self.Stop:
				for sub := range self.Subscribers {
					delete(self.Subscribers, sub)
					close(sub)
				}
			}
		}
	}()
}

func (self *ChanBroker) RegSubscriber() Subscriber {
	sub := make(Subscriber)
	self.RegSub <- sub
	return sub
}

func (self *ChanBroker) UnRegSubscriber(sub Subscriber) {
	self.UnRegSub <- sub
}

func (self *ChanBroker) StopPublish() {
	self.Stop <- true
}

func (self *ChanBroker) PubContent(c Content) {
	self.Contents <- c
}

type event struct {
	id   int
	info string
}

func main() {
	b := NewChanBroker()

	sub1 := b.RegSubscriber()
	sub2 := b.RegSubscriber()
	sub3 := b.RegSubscriber()
	go func(sub1 Subscriber) {
		for c := range sub1 {
			switch t := c.(type) {
			case string:
				fmt.Println(sub1, "string:", t)
			case int:
				fmt.Println(sub1, "int:", t)
			default:

			}
		}
	}(sub1)

	go func(sub2 Subscriber) {
		for c := range sub2 {
			switch t := c.(type) {
			case string:
				fmt.Println(sub2, "string:", t)
			case int:
				fmt.Println(sub2, "int:", t)
			case event:
				fmt.Println(sub2, "event:", t)
			default:

			}
		}
	}(sub2)

	go func(sub3 Subscriber) {
		c := <-sub3
		switch t := c.(type) {
		case string:
			fmt.Println(sub2, "string:", t)
		case int:
			fmt.Println(sub2, "int:", t)
		case event:
			fmt.Println(sub2, "event:", t)
		default:

		}

	}(sub3)

	ticker := time.NewTicker(time.Second)
	go func() {
		var prefix string = "pub_"
		i := 0
		for range ticker.C {
			i++
			if i%3 == 0 {
				var temp string
				temp = prefix + strconv.Itoa(i)
				b.PubContent(temp)
			} else if i%3 == 1 {
				b.PubContent(i)
			} else {
				ev := event{i, "event"}
				b.PubContent(ev)
			}

		}
	}()

	<-time.After(15 * time.Second)
}
