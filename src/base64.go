package main

import b64 "encoding/base64"
import "fmt"

func main() {
	data := "abc123456"

	src := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(src)

	dec, _ := b64.StdEncoding.DecodeString(src)
	fmt.Println(string(dec))
	fmt.Println()

	usrc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(usrc)
	udec, _ := b64.URLEncoding.DecodeString(usrc)
	fmt.Println(string(udec))

}
