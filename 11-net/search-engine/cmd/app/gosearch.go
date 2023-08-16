package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sort"

	"go-core-4/11-net/search-engine/pkg/crawler/spider"
	"go-core-4/11-net/search-engine/pkg/index"
)

const (
	godev       = "https://go.dev"
	practicalgo = "https://www.practical-go-lessons.com"
	docsFile    = "./docs.json"
)

func main() {
	searchString := "Go"
	var documents []index.Document

	f, err := os.Open(docsFile)
	if err != nil {
		log.Println(err)
		documents = search([]string{godev, practicalgo})
		addID(documents)
		sort.SliceStable(documents, func(i, j int) bool {
			return documents[i].ID < documents[j].ID
		})

		f, err = os.Create(docsFile)
		if err != nil {
			log.Println(err)
			return
		}
		defer f.Close()

		err = store(documents, f)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		documents, err = get(f)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(len(documents))
	}

	index := index.Make(documents)

	for _, id := range index[searchString] {
		i := sort.Search(len(documents), func(i int) bool {
			return documents[i].ID >= id
		})
		if i < len(documents) && documents[i].ID == id {
			fmt.Println(documents[i])
		}
	}
}

func search(urls []string) (res []index.Document) {
	for _, url := range urls {
		docs, err := scan(url)
		if err != nil {
			continue
		}
		res = append(res, docs...)
	}
	return res
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

func store(docs []index.Document, w io.Writer) error {
	b, err := json.Marshal(docs)
	if err != nil {
		return err
	}

	_, err = w.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func get(r io.Reader) (docs []index.Document, err error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &docs)
	if err != nil {
		return nil, err
	}
	return docs, nil
}
