package index

import (
	"sort"
	"strings"
)

// инвертированный интдекс: ключ - слово из Title, значение - слайс ID
type Index map[string][]int

// Document - документ, веб-страница, полученная поисковым роботом.
type Document struct {
	ID    int
	URL   string
	Title string
}

var index = make(map[string][]int)

// создание индекса
func Make(documents []Document) {
	for _, doc := range documents {
		words := strings.Split(doc.Title, " ")
		for _, word := range words {
			index[word] = append(index[word], doc.ID)
		}
	}
}

// поиск строки в отсканированных документах
func Search(s string, documents []Document) (result []string) {
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
