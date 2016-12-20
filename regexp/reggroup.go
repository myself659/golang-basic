package main

import "fmt"
import "regexp"

func main() {
	str := "123456"
	//rep := regexp.MustCompile(`1(.)3(.)5(.)`)
	rep := regexp.MustCompile("1(.)3(.)5(.)")
	str = rep.ReplaceAllString(str, "1($1)3($2)5($3)")

	fmt.Println(str) // => "1(2)3(4)5(6)"
	fmt.Println(rep.FindStringSubmatch(str))
	re := regexp.MustCompile("a(x*)b(y|z)c")
	fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbyc-"))
	fmt.Printf("%q\n", re.FindStringSubmatch("-abzc-"))
}
