package main

import (
	"github.com/bitly/go-nsq"
	"log"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("write_test", "ch", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("get a message:%#v", message)
		wg.Done()

		return nil
	}))
	err := q.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Panic("cannot connect")

	}
	wg.Wait()

}

/*

root@ubuntu:/share/go/src# go  run nsq-consumer.go
2016/07/18 23:56:30 INF    1 [write_test/ch] (127.0.0.1:4150) connecting to nsqd
2016/07/18 23:56:30 get a message:&nsq.Message{
ID:nsq.MessageID{0x30, 0x36, 0x38, 0x35, 0x35, 0x32, 0x30, 0x62, 0x33, 0x36, 0x66, 0x38, 0x35, 0x30, 0x30, 0x32},
Body:[]uint8{0x74, 0x65, 0x73, 0x74},
Timestamp:1468911376342020233,
Attempts:0x1,
 NSQDAddress:"127.0.0.1:4150",
 Delegate:(*nsq.connMessageDelegate)(0xc820024010),
 autoResponseDisabled:0,
 responded:0
 }

*/

/*
nsq高可用性

*/
