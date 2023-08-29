package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go-core-4/13-api/search-engine/pkg/index"
)

func (api *API) docs(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(index.Documents)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
