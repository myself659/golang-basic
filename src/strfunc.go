package main

import s "strings"
import "fmt"

var p = fmt.Println

func main() {
	a := "abc"
	b := "abcdef"
	xqurl := "http://bj.lianjia.com/chengjiao/c1111027376748/"
	//fmt.Println(s.Split(xqurl, "/"))
	aele := s.Split(xqurl, "/")
	for k, v := range aele {
		fmt.Println(k, v)
	}
	p(s.TrimPrefix(b, a))
	a3s := s.Split("南 北 | 其他 | 有电梯", "|")
	fmt.Println(a3s[0])
	fmt.Println(a3s[1])
	fmt.Println(a3s[2])
	p("Contains:", s.Contains("cloud compute", "cloud"))
	p("Count:", s.Count("cloud compute", "c"))
	p("HasPrefix:", s.HasPrefix("cloud", "c"))
	p("HasSuffix:", s.HasSuffix("Docker", "er"))
	p("HasSuffix:", s.HasSuffix("100.01亿", "亿"))
	p("Split:", s.Split("100.01亿", "亿"))
	as := s.Split(":100.01亿", ":")
	fmt.Printf("as:%v", as)
	p("Index:", s.Index("cloud compute", "c")) //取第一个
	p("Join:", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:", s.Repeat("ab", 5))
	p("Split:", s.Split("cloud-compute", "-"))
	p("ToLower:", s.ToLower("HELLO"))
	p("ToUpper:", s.ToTitle("hello"))

	p()
	p("Len:", len("hello"))
	p("char;", "hello"[1])
}
