package main

import "bytes"
import "fmt"
import "regexp"

func main() {
	// 简单正则匹配
	match, what := regexp.MatchString("p([a-z]+)ch", "peach")

	fmt.Println(match)
	fmt.Println(what)

	match0, what0 := regexp.MatchString("p([a-z]+)ch", "peat")

	fmt.Println(match0)
	fmt.Println(what0)

	// 编译成正则表达式，进行各个正则操作
	r, _ := regexp.Compile("p([a-z]+)(ch)")

	fmt.Println(r.MatchString("peach"))
	// 只匹配第一个匹配项 peach  leftmost match
	fmt.Println(r.FindString("peach  punch"))
	// 打印[0 5]
	fmt.Println(r.FindStringIndex("peach  punch"))

	//输出[peach ea] 子匹配项 [peach ea ch]
	fmt.Println(r.FindStringSubmatch("peach punch"))
	// 输出 [0 5 1 3] [0 5 1 3 3 5]
	fmt.Println(r.FindStringSubmatchIndex("peach  punch"))

	fmt.Println(r.FindAllString("peach punch  pinch", -1))

	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	// 只查找前两个匹配项，如果小于0 表示查找所有有匹配项
	fmt.Println(r.FindAllString("peach punch  pinch", 2))

	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach and  a big peach", "<fruit>"))
	// 只替换第一个匹配项
	//fmt.Println(r.ReplaceString("a peach and  a big peach", "<fruit>"))

	in := []byte("a  peach ")
	out := r.ReplaceAllFunc(in, bytes.ToUpper) //对匹配项按传入函数操作
	fmt.Println(string(out))

}

/* 同python进行比较 */
