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
	readOK := false

	if _, err := os.Stat(docsFile); err == nil {
		documents, err = readFromFile()
		if err == nil {
			readOK = true
		} else {
			log.Println(err)
		}
	}

	if !readOK {
		documents = scan([]string{godev, practicalgo})
		sort.SliceStable(documents, func(i, j int) bool {
			return documents[i].ID < documents[j].ID
		})
		saveToFile(documents)
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

func readFromFile() (documents []index.Document, err error) {
	f, err := os.Open(docsFile)
	if err != nil {
		return nil, err
	}

	documents, err = get(f)
	if err != nil {
		return nil, err
	}

	return documents, nil
}

func scan(urls []string) (data []index.Document) {
	s := spider.New()

	for _, url := range urls {
		docs, err := s.Scan(url, 2)

		if err != nil {
			log.Println(err)
			continue
		}

		len := len(data)

		for i, doc := range docs {
			data = append(data, index.Document{ID: len + i, URL: doc.URL, Title: doc.Title})
		}
	}

	return data
}

func saveToFile(docs []index.Document) {
	f, err := os.Create(docsFile)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	b, err := json.Marshal(docs)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = f.Write(b)
	if err != nil {
		log.Println(err)
		return
	}
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
