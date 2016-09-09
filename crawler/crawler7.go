package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/jackdanger/collectlinks"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var visited = make(map[string]bool)

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

	links := collectlinks.All(resp.Body)

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
