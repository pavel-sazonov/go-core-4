package scanner

import (
	"go-core-4/02hw/search-engine/pkg/crawler/spider"
)

// возвращает найденные на страницах ссылки
func URLs(urls []string) []string {
	data := make([]string, 0)

	for _, url := range urls {
		urls, err := scan(url)
		if err != nil {
			continue
		}
		data = append(data, urls...)
	}
	return data
}

// возвращает найденные на странице ссылки
func scan(url string) (urls []string, err error) {
	s := spider.New()
	res := make([]string, 0)
	docs, err := s.Scan(url, 2)
	if err != nil {
		return nil, err
	}
	for _, doc := range docs {
		res = append(res, doc.URL)
	}
	return res, nil
}
