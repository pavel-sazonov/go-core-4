package main

import (
	"log"
	"sync"

	"go-core-4/13-api/search-engine/pkg/index"
	"go-core-4/13-api/search-engine/pkg/netsrv"
	"go-core-4/13-api/search-engine/pkg/webapp"
)

const (
	godev       = "https://go.dev"
	practicalgo = "https://www.practical-go-lessons.com"
)

func main() {
	getData()
	index.MakeIndex()

	var wg sync.WaitGroup

	go func() {
		defer wg.Done()
		err := webapp.StartServer()
		if err != nil {
			log.Println(err)
		}
	}()

	err := netsrv.Start()
	if err != nil {
		log.Println(err)
	}

	wg.Wait()
}

func getData() {
	if index.IsFileExist() {
		err := index.ReadFromFile()
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	index.Scan([]string{godev, practicalgo})
	err := index.SaveToFile()
	if err != nil {
		log.Fatal(err)
	}
}
