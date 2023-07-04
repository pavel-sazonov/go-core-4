package index

import (
	"strings"
)

// инветрированный интдекс: ключ - слово из Title, значение - ID
type Index map[string][]int

// Document - документ, веб-страница, полученная поисковым роботом.
type Document struct {
	ID    int
	URL   string
	Title string
}

func Make(documents []Document) Index {
	var index = make(map[string][]int)

	for _, doc := range documents {
		words := strings.Split(doc.Title, " ")
		for _, word := range words {
			index[word] = append(index[word], doc.ID)
		}
	}
	return index
}
