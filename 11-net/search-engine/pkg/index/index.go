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
type index map[string][]int

// Document - документ, веб-страница, полученная поисковым роботом.
type Document struct {
	ID    int
	URL   string
	Title string
}

var Documents []Document
var Index = make(map[string][]int)

// создание индекса
func Make() {
	for _, doc := range Documents {
		words := strings.Split(doc.Title, " ")
		for _, word := range words {
			Index[word] = append(Index[word], doc.ID)
		}
	}
}

// поиск строки в отсканированных документах
func Search(s string) (result []string) {
	for _, id := range Index[s] {
		i := sort.Search(len(Documents), func(i int) bool {
			return Documents[i].ID >= id
		})
		if i < len(Documents) && Documents[i].ID == id {
			result = append(result, Documents[i].Title)
		}
	}
	return result
}

// чтение отсканированных документов из сохраненного файла
// при отсутствии файла, сканирование и сохранение результатов в новый файл
func ReadOrScanDocuments() {
	readOK := false

	if _, err := os.Stat(docsFile); err == nil {
		Documents, err = readFromFile()
		if err == nil {
			readOK = true
		} else {
			log.Println(err)
		}
	}

	if !readOK {
		Documents = scan([]string{godev, practicalgo})
		sort.SliceStable(Documents, func(i, j int) bool {
			return Documents[i].ID < Documents[j].ID
		})
		saveToFile(Documents)
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
