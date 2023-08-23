package main

import (
	"sync"

	"go-core-4/11-net/search-engine/pkg/index"
	"go-core-4/11-net/search-engine/pkg/netsrv"
	"go-core-4/11-net/search-engine/pkg/webapp"
)

func main() {
	var wg sync.WaitGroup
	index.GetDocuments()

	wg.Add(1)
	go netsrv.Start(&wg)

	wg.Add(1)
	go webapp.StartServer(&wg)

	wg.Wait()
}
