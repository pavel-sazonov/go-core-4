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
	res := scanner.Result([]string{godev, practicalgo})
	for k, v := range res {
		fmt.Println(k, v)
	}

}
