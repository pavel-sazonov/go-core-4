package main

import (
	"go-core-4/11-net/search-engine/pkg/index"
	"go-core-4/11-net/search-engine/pkg/webapp"
)

func main() {
	index.ReadOrScanDocuments()
	index.Make()

	// netsrv.Start()
	webapp.StartServer()
}
