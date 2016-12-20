package main

import (
	"bytes"
	_ "encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/url"
	_ "os"
)

func main() {

	url := "http://192.168.80.141:10601/api/w/invoke.htm?_s=notice&_m=addClassChatRule"
	var b = []byte("abc")
	body := bytes.NewBuffer([]byte(b))
	res, err := http.Post(url, "application/json;charset=utf-8", body)
	if err != nil {
		log.Fatal(err)
		return
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s", result)
}
