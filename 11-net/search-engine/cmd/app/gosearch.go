package main

import (
	"go-core-4/11-net/search-engine/pkg/index"
	"go-core-4/11-net/search-engine/pkg/netsrv"
)

func main() {
	index.GetDocuments()
	index.Make()

	netsrv.Start()
}
