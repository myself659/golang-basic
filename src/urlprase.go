package main

import "fmt"
import "net/url"
import "strings"

func main() {

	s := "https://www.google.co.jp/?gfe_rd=cr&ei=nJUpV9B4ssryB5uribAB&gws_rd=ssl#newwindow=1&safe=off&q=game+of+thrones+season+6+episode+3"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)

	/*
		fmt.Println(u.User)
		fmt.Println(u.User.Username())

		p, _ := u.User.Password()
		fmt.Println(p)
	*/

	fmt.Println(u.Host)
	h := strings.Split(u.Host, ":")
	fmt.Println(h[0])
	//fmt.Println(h[1])

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)

	m, _ := url.ParseQuery(u.RawQuery)

	fmt.Println(m)

}

/* url 地址格式 */
