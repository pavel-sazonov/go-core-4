package index

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sort"
	"strings"

	"go-core-4/11-net/search-engine/pkg/crawler/spider"
)

const (
	docsFile    = "./docs.json"
	godev       = "https://go.dev"
	practicalgo = "https://www.practical-go-lessons.com"
)

// инвертированный интдекс: ключ - слово из Title, значение - слайс ID
type Index map[string][]int

// Document - документ, веб-страница, полученная поисковым роботом.
type Document struct {
	ID    int
	URL   string
	Title string
}

var documents []Document
var index = make(map[string][]int)

// создание индекса
func Make() {
	for _, doc := range documents {
		words := strings.Split(doc.Title, " ")
		for _, word := range words {
			index[word] = append(index[word], doc.ID)
		}
	}
}

// поиск строки в отсканированных документах
func Search(s string) (result []string) {
	for _, id := range index[s] {
		i := sort.Search(len(documents), func(i int) bool {
			return documents[i].ID >= id
		})
		if i < len(documents) && documents[i].ID == id {
			result = append(result, documents[i].Title)
		}
	}
	return result
}

// чтение отсканированных документов из сохраненного файла
// при отсутствии файла, сканирование и сохранение результатов в новый файл
func ReadOrScanDocuments() {
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
}

// получение результатов сканирования веб страниц из файла
func readFromFile() (documents []Document, err error) {
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
func scan(urls []string) (data []Document) {
	s := spider.New()

	for _, url := range urls {
		docs, err := s.Scan(url, 2)

		if err != nil {
			log.Println(err)
			continue
		}

		len := len(data)

		for i, doc := range docs {
			data = append(data, Document{ID: len + i, URL: doc.URL, Title: doc.Title})
		}
	}

	return data
}

// сохранение результата сканирования веб страниц в файл
func saveToFile(docs []Document) {
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
