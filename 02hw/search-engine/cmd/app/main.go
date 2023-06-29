package main

import (
	"fmt"
	"go-core-4/02hw/search-engine/pkg/scanner"
)

const (
	godev       = "https://go.dev"
	practicalgo = "https://www.practical-go-lessons.com"
)

func main() {
	urls := scanner.URLs([]string{godev, practicalgo})
	for i, u := range urls {
		fmt.Println(i, ":", u)
	}

}
