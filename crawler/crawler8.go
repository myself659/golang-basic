package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	// "github.com/jackdanger/collectlinks"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var visited = make(map[string]bool)

func GetAll(httpBody io.Reader) []string {
	links := []string{}
	col := []string{}
	page := html.NewTokenizer(httpBody)
	//fmt.Println("Tokenizer:", page)
	fmt.Printf("Tokenizer:%+v\n", page)
	for {
		tokenType := page.Next()
		if tokenType == html.ErrorToken {
			return links // 解析失败或者结束，返回链接
		}
		token := page.Token()
		//fmt.Println("Token:", token)
		fmt.Printf("Token:%+v\n", token)
		if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					tl := trimHash(attr.Val) // 对应内容
					col = append(col, tl)
					resolv(&links, col)
				}
			}
		}
	}
}

// trimHash slices a hash # from the link
func trimHash(l string) string {
	fmt.Println("trimHash:", l)
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

func filter(uri string) bool {
	return strings.Contains(uri, "myself659.github.io")
}

func enqueue(uri string, queue chan string) {
	fmt.Println("fetching", uri)
	visited[uri] = true
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	client := http.Client{Transport: transport}
	resp, err := client.Get(uri)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	links := GetAll(resp.Body)

	for _, link := range links {
		absolute := fixUrl(link, uri)
		if absolute != "" {
			if filter(uri) && !visited[absolute] {
				go func() {
					queue <- absolute
				}()
			}
		}
	}
}

func fixUrl(href string, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return ""
	}
	baseurl, err := url.Parse(base)
	if err != nil {
		return ""
	}
	uri = baseurl.ResolveReference(uri)

	return uri.String()
}

func main() {
	flag.Parse()

	args := flag.Args()
	fmt.Println(args)
	if len(args) < 1 {
		fmt.Println("Please specify start page")
		os.Exit(1)
	}

	queue := make(chan string)
	go func() {
		queue <- args[0]
	}()

	for uri := range queue {
		enqueue(uri, queue)
	}
}

/*
Token: <a href="http://blog.studygolang.com" target="_blank">
trimHash: http://blog.studygolang.com

*/

/*
type Token struct {
	Type     TokenType
	DataAtom atom.Atom
	Data     string
	Attr     []Attribute
}
*/

/*
Token: <a href="/topics/1894" title="[郑州] 招聘golang、数据仓库工程师、功能/性能测试，移动端leader，类比百度 T5 及以上。">
trimHash: /topics/1894
Token: [郑州] 招聘golang、数据仓库工程师、功能/性能测试，移动端leader，类比百度 T5 及以上。
*/

/*
<a href="/topics/1894#commentForm" title="查看评论">7 <span class="glyphicon glyphicon-comment"></span></a>
*/

/*
(gdb) p token.Attr.array[0]
$9 = {Namespace = 0x0 "", Key = 0x843010 "href", Val = 0xc82000b0e0 "http://docs.studygolang.com"}
(gdb) p token.Attr.array[1]
$10 = {Namespace = 0x0 "", Key = 0x842feb "target", Val = 0xc82000f744 "_blank"}
(gdb)
*/

/*
<li><a href="http://docs.studygolang.com" target="_blank">英文文档</a></li>
*/

/*
Token:<li>
Token:<a href="http://docs.studygolang.com" target="_blank">
Token:英文文档
Token:</a>
Token:</li>
*/

/*
<a href="/topics/1894" title="[郑州] 招聘golang、数据仓库工程师、功能/性能测试，移动端leader，类比百度 T5 及以上。">
[郑州] 招聘golang、数据仓库工程师、功能/性能测试，移动端leader，类比百度 T5 及以上。</a>
*/

/*
Token:<a href="/topics/1894" title="[郑州] 招聘golang、数据仓库工程师、功能/性能测试，移动端leader，类比百度 T5 及以上。">

Token:[郑州] 招聘golang、数据仓库工程师、功能/性能测试，移动端leader，类比百度 T5 及以上。
Token:</a>
*/

/*
(gdb) p token
$11 = {Type = 2, DataAtom = 1, Data = 0x840e60 "a", Attr = {array = 0xc8201313e0, len = 2, cap = 2}}
(gdb) p token->Attr.array[0]
A syntax error in expression, near `>Attr.array[0]'.
(gdb) p token.Attr.array[0]
$12 = {Namespace = 0x0 "", Key = 0x843010 "href", Val = 0xc82015e6e0 "/topics/1894"}
(gdb) p token.Attr.array[1]
$13 = {Namespace = 0x0 "", Key = 0x842e41 "title",
  Val = 0xc820126500 "[郑州] 招聘golang、数据仓库工程师、功能/性能测试，移动端leader，类比百度 T5 及以上。"}
(gdb)
*/
