package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Reminder struct {
	Id        int64     `json:"id"`
	Message   string    `sql:"size:1024" json:"message"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"-"`
}

var method = flag.String("method", "POST", "GET POST ")
var msg = flag.String("msg", "defalt msg", "just a string")

func pro() {
	fmt.Println(*method)
	fmt.Println(*msg)
	switch *method {
	case "POST":

		temp := &Reminder{
			Id:        20,
			Message:   *msg,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		enc, err := json.Marshal(temp)
		if err != nil {
			fmt.Println("encoding json ", err)
		}
		fmt.Println(string(enc))
		ioread := strings.NewReader(string(enc))
		//encoding to json
		resp, err := http.Post("http://127.0.0.1:8008/reminders", "application/json", ioread)
		fmt.Println(resp.Status)
		if err !=nil{
		 fmt.Println("post", err)
		}
	case "GET":
		resp, err := http.Get("http://127.0.0.1:8008/reminders")

		fmt.Println(resp, err)
		fmt.Println(resp.Status)
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Proto)
		fmt.Println(resp.ProtoMajor)
		fmt.Println(resp.ProtoMinor)
		fmt.Println(resp.Header)
		fmt.Println(resp.ContentLength)
		p := make([]byte, resp.ContentLength)
		resp.Body.Read(p)
		fmt.Println(string(p))
		//fmt.Println(resp.Body)

	}
}
func main() {
	flag.Parse()
	pro()

}
