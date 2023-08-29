package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"

	"go-core-4/13-api/search-engine/pkg/index"
)

func (api *API) doc(w http.ResponseWriter, r *http.Request) {
	search := mux.Vars(r)["search"]
	doc := index.Search(search)

	if len(doc) == 0 {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	err := json.NewEncoder(w).Encode(doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (api *API) docs(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(index.Documents)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
