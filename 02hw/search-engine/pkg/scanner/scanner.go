package scanner

import (
	"go-core-4/02hw/search-engine/pkg/crawler/spider"
)

func Result(urls []string) map[string][]string {
	m := make(map[string][]string)

	for _, url := range urls {
		urls, err := scan(url)
		if err != nil {
			m[url] = make([]string, 0)
			continue
		}
		m[url] = urls
	}
	return m
}

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
