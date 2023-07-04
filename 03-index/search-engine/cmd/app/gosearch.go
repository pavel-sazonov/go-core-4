package main

import (
	"flag"
	"fmt"
	"go-core-4/03-index/search-engine/pkg/crawler/spider"
	"go-core-4/03-index/search-engine/pkg/index"
	"strings"
)

const (
	godev       = "https://go.dev"
	practicalgo = "https://www.practical-go-lessons.com"
)

var s = flag.String("s", "", "search argument")

func main() {
	flag.Parse()
	documents := scanResult([]string{godev, practicalgo})
	index := index.Make(documents)

	for _, v := range search(*s, documents) {
		fmt.Println(v)
	}
}

// поиск строки в слайсе Document
func search(s string, documents []index.Document) (result []string) {
	for _, doc := range documents {
		if strings.Contains(doc.URL, s) {
			result = append(result, doc.URL)
		}
	}

	return result
}

func scanResult(urls []string) (documents []index.Document) {
	for _, url := range urls {
		docs, err := scan(url)
		if err != nil {
			continue
		}
		documents = append(documents, docs...)
	}
	return documents
}

func scan(url string) (documents []index.Document, err error) {
	s := spider.New()
	docs, err := s.Scan(url, 2)
	if err != nil {
		return nil, err
	}
	for _, doc := range docs {
		documents = append(documents, index.Document{ID: doc.ID, URL: doc.URL, Title: doc.Title})
	}
	return documents, nil
}
