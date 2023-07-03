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

// Document - документ, веб-страница, полученная поисковым роботом.
type Document struct {
	ID    int
	URL   string
	Title string
}

func main() {
	flag.Parse()
	documents := scanResult([]string{godev, practicalgo})

	for _, v := range search(*s, documents) {
		fmt.Println(v)
	}
}

// поиск строки в слайсе Document
func search(s string, documents []Document) (result []string) {
	for _, doc := range documents {
		if strings.Contains(doc.URL, s) {
			result = append(result, doc.URL)
		}
	}

	return result
}

func scanResult(urls []string) (documents []Document) {
	for _, url := range urls {
		docs, err := scan(url)
		if err != nil {
			continue
		}
		documents = append(documents, docs...)
	}
	return documents
}

func scan(url string) (documents []Document, err error) {
	s := spider.New()
	docs, err := s.Scan(url, 2)
	if err != nil {
		return nil, err
	}
	for _, doc := range docs {
		documents = append(documents, Document{doc.ID, doc.URL, doc.Title})
	}
	return documents, nil
}
