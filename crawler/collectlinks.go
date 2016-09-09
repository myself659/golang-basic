// Package collectlinks does extraordinarily simple operation of parsing a given piece of html
// and providing you with all the hyperlinks hrefs it finds.
package collectlinks

import (
	"io"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

// All takes a reader object (like the one returned from http.Get())
// It returns a slice of strings representing the "href" attributes from
// anchor links found in the provided html.
// It does not close the reader passed to it.
/*

一个链接实例 <a href="/im-message/" title="IM后端系统设计总结(1)">IM后端系统设计总结(1)</a>
分析一个HTML网页中链接

*/
func All(httpBody io.Reader) []string {
	links := []string{}
	col := []string{}
	page := html.NewTokenizer(httpBody)
	for {
		tokenType := page.Next()
		if tokenType == html.ErrorToken {
			return links // 解析失败或者结束，返回链接
		}
		token := page.Token()
		if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					tl := trimHash(attr.Val)
					col = append(col, tl)
					resolv(&links, col)
				}
			}
		}
	}
}

// trimHash slices a hash # from the link
func trimHash(l string) string {
	if strings.Contains(l, "#") { // html语言#表示什么，#表示网页的其他部分用于页面跳转
		var index int
		for n, str := range l {
			if strconv.QuoteRune(str) == "'#'" {
				index = n
				break
			}
		}
		return l[:index]
	}
	return l
}

// check looks to see if a url exits in the slice.
func check(sl []string, s string) bool {
	var check bool
	for _, str := range sl {
		if str == s {
			check = true
			break
		}
	}
	return check
}

// resolv adds links to the link slice and insures that there is no repetition
// in our collection.
func resolv(sl *[]string, ml []string) {
	for _, str := range ml {
		if check(*sl, str) == false { // 如果重复，则不添加
			*sl = append(*sl, str)
		}
	}
}
