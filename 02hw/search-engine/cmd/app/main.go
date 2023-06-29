package main

import (
	"fmt"
	"go-core-4/02hw/search-engine/pkg/scanner"
	"os"
	"strings"
)

const (
	godev       = "https://go.dev"
	practicalgo = "https://www.practical-go-lessons.com"
	gosearch    = "gosearch"
)

func main() {
	c, err := commands(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}

	urls := scanner.URLs([]string{godev, practicalgo})
	for i, u := range urls {
		fmt.Println(i, ":", u)
	}

	for _, s := range search(c, urls) {
		fmt.Println(s)
	}

}

func commands(s []string) ([]string, error) {
	if s[0] != gosearch {
		return nil, fmt.Errorf("команда не найдена")
	}

	res := make([]string, 0)

	for _, c := range s[2:] {
		res = append(res, c)
	}
	return res, nil
}

func search(s []string, urls []string) []string {
	res := make([]string, 0)

	for _, c := range s {
		for _, url := range urls {
			if strings.Contains(url, c) {
				res = append(res, url)
			}
		}
	}

	return res
}
