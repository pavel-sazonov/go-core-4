package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/exp/slices"

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

func (api *API) deleteDoc(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	i := index.IndexByID(id)
	if len(index.Documents) > i && index.Documents[i].ID == id {
		err = json.NewEncoder(w).Encode(index.Documents[i])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		index.Documents = slices.Delete(index.Documents, i, i+1)
	} else {
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}
