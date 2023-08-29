package index

import (
	"encoding/json"
	"io"
	"os"
	"sort"
	"strings"

	"go-core-4/13-api/search-engine/pkg/crawler/spider"
)

const (
	docsFile = "./docs.json"
)

// инвертированный интдекс: ключ - слово из Title, значение - слайс ID
type Index map[string][]int

// Document - документ, веб-страница, полученная поисковым роботом.
type Document struct {
	ID    int
	URL   string
	Title string
}

var index Index
var Documents []Document

func GetIndex() Index {
	return index
}

// поиск строки в отсканированных документах
func Search(s string) (result []string) {
	for _, id := range index[s] {
		i := sort.Search(len(Documents), func(i int) bool {
			return Documents[i].ID >= id
		})
		if i < len(Documents) && Documents[i].ID == id {
			result = append(result, Documents[i].Title)
		}
	}
	return result
}

// проверка, существует ли файл с данным отсканированных веб страниц
func IsFileExist() bool {
	_, err := os.Stat(docsFile)
	return err == nil
}

// получение результатов сканирования веб страниц из файла
func ReadFromFile() error {
	f, err := os.Open(docsFile)
	if err != nil {
		return err
	}

	data, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &Documents)
	if err != nil {
		return err
	}
	return nil
}

// сканирование веб страниц по ссылкам
func Scan(urls []string) {
	s := spider.New()
	data := make([]Document, 0)

	for _, url := range urls {
		docs, err := s.Scan(url, 2)

		if err != nil {
			continue
		}

		len := 0

		for i, doc := range docs {
			data = append(data, Document{ID: len + i, URL: doc.URL, Title: doc.Title})
		}
	}

	Documents = data
}

// сохранение результата сканирования веб страниц в файл
func SaveToFile() error {
	f, err := os.Create(docsFile)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := json.Marshal(Documents)
	if err != nil {
		return err
	}

	_, err = f.Write(b)
	if err != nil {
		return err
	}

	return nil
}

// создание индекса
func MakeIndex() {
	index = make(map[string][]int)

	for _, doc := range Documents {
		words := strings.Split(doc.Title, " ")
		for _, word := range words {
			index[word] = append(index[word], doc.ID)
		}
	}
}
