package main

import "fmt"
import "regexp"

func main() {
	r := `(\S+):       # 名前
         ([\d\-]+)    # 電話番号
    `
	str := "hoge:0045-111-2222 boke:0045-222-2222"
	result := find_x(r, str)
	fmt.Println(result[0]) // => "[hoge:0045-111-2222 hoge 0045-111-2222]"
	fmt.Println(result[1]) // => "[boke:0045-222-2222 boke 0045-222-2222]"
}

func find_x(r, str string) [][]string {

	com := regexp.MustCompile(`(?m)(\s+)|(\#.*$)`)
	r = com.ReplaceAllString(r, "")
	reg := regexp.MustCompile(r)
	return reg.FindAllStringSubmatch(str, -1)
}
