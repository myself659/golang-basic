package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := "aa{abc},  {def}, {ef}"
	r, _ := regexp.Compile(`(?U){.*}`)
	ar := r.FindAllStringSubmatch(s, -1)
	fmt.Println(ar)

}
