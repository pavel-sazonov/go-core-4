package scanner

import (
	"go-core-4/02hw/search-engine/pkg/crawler/spider"
)

// возвращает найденные на страницах ссылки
func URLs(urls []string) []string {
	var data []string

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
	var res []string
	docs, err := s.Scan(url, 2)
	if err != nil {
		return nil, err
	}
	for _, doc := range docs {
		res = append(res, doc.URL)
	}
	return res, nil
}
