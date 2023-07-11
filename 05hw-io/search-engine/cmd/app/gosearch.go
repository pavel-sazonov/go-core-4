package main

import (
	"flag"
	"fmt"
	"go-core-4/05hw-io/search-engine/pkg/crawler/spider"
	"go-core-4/05hw-io/search-engine/pkg/index"
	"sort"
)

const (
	godev       = "https://go.dev"
	practicalgo = "https://www.practical-go-lessons.com"
)

var s = flag.String("s", "", "search argument")

func main() {
	flag.Parse()
	urls := []string{godev, practicalgo}
	var documents []index.Document

	for _, url := range urls {
		docs, err := scan(url)
		if err != nil {
			continue
		}
		documents = append(documents, docs...)
	}

	addID(documents)

	sort.SliceStable(documents, func(i, j int) bool {
		return documents[i].ID < documents[j].ID
	})

	index := index.Make(documents)

	for _, id := range index[*s] {
		i := sort.Search(len(documents), func(i int) bool {
			return documents[i].ID >= id
		})
		if i < len(documents) && documents[i].ID == id {
			fmt.Println(documents[i])
		}
	}
}

func scan(url string) (documents []index.Document, err error) {
	s := spider.New()
	docs, err := s.Scan(url, 2)
	if err != nil {
		return nil, err
	}
	for _, doc := range docs {
		documents = append(documents, index.Document{URL: doc.URL, Title: doc.Title})
	}
	return documents, nil
}

func addID(docs []index.Document) {
	for i := range docs {
		docs[i].ID = i
	}
}
