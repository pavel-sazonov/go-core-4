package main

import (
	"log"
	"sync"

	"go-core-4/11-net/search-engine/pkg/index"
	"go-core-4/11-net/search-engine/pkg/netsrv"
	"go-core-4/11-net/search-engine/pkg/webapp"
)

const (
	godev       = "https://go.dev"
	practicalgo = "https://www.practical-go-lessons.com"
)

func main() {
	getData()
	index.MakeIndex()

	var wg sync.WaitGroup

	wg.Add(1)
	go webapp.StartServer(&wg)

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
