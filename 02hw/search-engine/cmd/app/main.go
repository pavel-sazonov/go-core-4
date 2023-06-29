package main

import (
	"flag"
	"fmt"
	"go-core-4/02hw/search-engine/pkg/scanner"
	"strings"
)

const (
	godev       = "https://go.dev"
	practicalgo = "https://www.practical-go-lessons.com"
	gosearch    = "gosearch"
)

var s = flag.String("s", "", "search argument")

func main() {
	urls := scanner.URLs([]string{godev, practicalgo})
	for i, u := range urls {
		fmt.Println(i, ":", u)
	}

	for _, s := range search(*s, urls) {
		fmt.Println(s)
	}

}

func search(s string, urls []string) []string {
	res := make([]string, 0)

	for _, url := range urls {
		if strings.Contains(url, s) {
			res = append(res, url)
		}
	}

	return res
}
