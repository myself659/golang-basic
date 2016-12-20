package main

import (
	"fmt"
	"github.com/nilslice/email"
)

func main() {
	msg := email.Message{
		To:      "xx@qq.com",         // do not add < > or name in quotes
		From:    "Name <xx@126.com>", // ok to format in From field
		Subject: "A simple email",
		Body:    "Plain text email body. HTML not yet supported, but send a PR!",
	}

	err := msg.Send()
	if err != nil {
		fmt.Println(err)
	}
}
