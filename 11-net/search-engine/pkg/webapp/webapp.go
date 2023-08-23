package webapp

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"

	"go-core-4/11-net/search-engine/pkg/index"
)

func StartServer(wg *sync.WaitGroup) {
	defer wg.Done()

	const addr = ":8080"
	mux := mux.NewRouter()

	// Регистрация обработчика для URL `/` в маршрутизаторе по умолчанию.
	mux.HandleFunc("/{name}", mainHandler).Methods(http.MethodGet)

	srv := &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		Handler:      mux,
		Addr:         addr,
	}

	// Старт сетевой службы веб-сервера.
	listener, err := net.Listen("tcp4", addr)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(srv.Serve(listener))
}

// HTTP-обработчик
func mainHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if v, ok := vars["name"]; ok {
		data, err := data(v)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func data(name string) (data []byte, err error) {
	switch name {
	case "index":
		data, err = json.Marshal(index.Index)
	case "docs":
		data, err = json.Marshal(index.Documents)
	}

	if err != nil {
		return nil, err
	}

	return data, nil
}
