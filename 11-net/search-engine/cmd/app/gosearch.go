package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sort"

	"go-core-4/11-net/search-engine/pkg/crawler/spider"
	"go-core-4/11-net/search-engine/pkg/index"
	"go-core-4/11-net/search-engine/pkg/netsrv"
)

const (
	godev       = "https://go.dev"
	practicalgo = "https://www.practical-go-lessons.com"
	docsFile    = "./docs.json"
)

func main() {
	documents := readOrScan()
	index.Make(documents)

	netsrv.Start(documents)
}

// чтение отсканированных документов из сохраненного файла
// при отсутствии файла, сканирование и сохранение результатов в новый файл
func readOrScan() (data []index.Document) {
	readOK := false

	if _, err := os.Stat(docsFile); err == nil {
		data, err = readFromFile()
		if err == nil {
			readOK = true
		} else {
			log.Println(err)
		}
	}

	if !readOK {
		data = scan([]string{godev, practicalgo})
		sort.SliceStable(data, func(i, j int) bool {
			return data[i].ID < data[j].ID
		})
		saveToFile(data)
	}

	return data
}

// получение результатов сканирования веб страниц из файла
func readFromFile() (documents []index.Document, err error) {
	f, err := os.Open(docsFile)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &documents)
	if err != nil {
		return nil, err
	}

	return documents, nil
}

// сканирование веб страниц по ссылкам
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

// сохранение результата сканирования веб страниц в файл
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
