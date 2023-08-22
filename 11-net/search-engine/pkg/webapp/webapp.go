package webapp

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func StartServer() {
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
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<html><body><h2>Hi, %v</h2></body></html>", vars["name"])
}
