package main

import (
	"flag"
	"fmt"
	"go-core-4/03-index/search-engine/pkg/crawler/spider"
	"strings"
)

const (
	godev       = "https://go.dev"
	practicalgo = "https://www.practical-go-lessons.com"
)

var s = flag.String("s", "", "search argument")

func main() {
	flag.Parse()
	urls := scanResult([]string{godev, practicalgo})

	for _, v := range search(*s, urls) {
		fmt.Println(v)
	}
}

// поиск строки в слайсе урлов
func search(s string, urls []string) []string {
	var res []string

	for _, url := range urls {
		if strings.Contains(url, s) {
			res = append(res, url)
		}
	}

	return res
}

func scanResult(urls []string) []string {
	var data []string

	for _, url := range urls {
		urls, err := scan(url)
		if err != nil {
			continue
		}
		data = append(data, urls...)
	}
	return data
}

func scan(url string) (urls []string, err error) {
	s := spider.New()
	var res []string
	docs, err := s.Scan(url, 2)
	if err != nil {
		return nil, err
	}
	for _, doc := range docs {
		res = append(res, doc.URL)
	}
	return res, nil
}
