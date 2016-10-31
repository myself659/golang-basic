package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("before log fatal")

	_, err := os.Open("abc")

	checkError("open failed", err)

	fmt.Println("after log fatal")

}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
