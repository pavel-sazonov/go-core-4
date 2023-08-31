package main

import (
	"log"
	"net/http"
	"sync"

	"go-core-4/13-api/search-engine/pkg/api"
	"go-core-4/13-api/search-engine/pkg/index"
	"go-core-4/13-api/search-engine/pkg/netsrv"
	"go-core-4/13-api/search-engine/pkg/webapp"
)

const (
	godev       = "https://go.dev"
	practicalgo = "https://www.practical-go-lessons.com"
)

type server struct {
	api *api.API
}

func main() {
	getData()
	index.MakeIndex()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := webapp.StartServer()
		if err != nil {
			log.Println(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		srv := new(server)

		srv.api = api.New()

		http.ListenAndServe(":8082", srv.api.Router())
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
