package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func main() {
	html_code := strings.NewReader(`
<html>
    <body>
    <ul class="listContent" log-mod="list">
    <div class="houseInfo">
	<span class="houseIcon"/>
	<a title="官书院网签" href="http://bj.lianjia.com/chengjiao/c1111027374707/">30天成交2套</a>
	<span class="cutLine">|</span>
	<a title="官书院租房" href="http://bj.lianjia.com/zufang/c1111027374707/">3套正在出租</a>
	</div>
	<div class="houseInfo">
	<span class="houseIcon"/>
	<a title="官书院网签" href="http://bj.lianjia.com/chengjiao/c1111027374707/">30天成交2套</a>
	<span class="cutLine">|</span>
	<a title="官书院租房" href="http://bj.lianjia.com/zufang/c1111027374707/">3套正在出租</a>
	</div>
    </ul>
    <div class="con">
	<a href="/chengjiao/yangpu/">杨浦</a>
	<a href="/chengjiao/huangxinggongyuan/">黄兴公园</a>
	
	<span>| </span>高区/6层
	
	
	<span>| </span>朝南
	
	
	<span>| </span>中装
										
	</div>
	<div class="con">

	</div>
        <h1>
            <span class="text title">Go </span>
        </h1>
        <p>
            <span class="text">totally </span>
            <span class="post">kicks </span>
        </p>
        <p>
            <span class="text">hacks </span>
        </p>
    </body>
<html>
    `)
	doc, _ := goquery.NewDocumentFromReader(html_code)
	//fmt.Println(doc.Find("ul.listContent").Text())
	listContent := doc.Find("ul.listContent")
	listContent.Find("span.houseIcon").Each(func(i int, s *goquery.Selection) {
		linka := s.Find("a")
		link, _ := linka.Attr("href")
		fmt.Println(link)
	})
	con := doc.Find("div.con")
	fmt.Println(con.Length())
	fmt.Println(con.Size())
	fmt.Println(con.Html())
	scon := con.Text()
	fmt.Println(scon)
	i := 0
	//获取长度 怎么出现了3次
	doc.Find(".text.title").Each(func(i int, s *goquery.Selection) {
		class, _ := s.Attr("class")

		fmt.Println(i, class, s.Text())
	})
	i++
	doc.Find(".title").Each(func(i int, s *goquery.Selection) {
		class, _ := s.Attr("class")
		fmt.Println(i, class, s.Text())
	})
	i++
	doc.Find("span.text titel").Each(func(i int, s *goquery.Selection) {
		class, _ := s.Attr("class")
		fmt.Println(i, class, s.Text())
	})
	i++
	// 这一个为什么没有match 出来
	doc.Find("span.text").Each(func(i int, s *goquery.Selection) {
		class, _ := s.Attr("class")
		fmt.Println(i, class, s.Text())
	})

}
