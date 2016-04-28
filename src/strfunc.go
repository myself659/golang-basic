package main

import s "strings"
import "fmt"

var p = fmt.Println

func main() {
	p("Contains:", s.Contains("cloud compute", "cloud"))
	p("Count:", s.Count("cloud compute", "c"))
	p("HasPrefix:", s.HasPrefix("cloud", "c"))
	p("HasSuffix:", s.HasSuffix("Docker", "er"))
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
