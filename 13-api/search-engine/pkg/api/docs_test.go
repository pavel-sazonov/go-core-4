package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/gorilla/mux"

	"go-core-4/13-api/search-engine/pkg/index"
)

var api *API

func TestMain(m *testing.M) {
	api = new(API)
	api.router = mux.NewRouter()
	api.endpoints()
	os.Exit(m.Run())
}

func TestAPI_docs(t *testing.T) {
	index.Documents = []index.Document{
		{ID: 1, URL: "url", Title: "title"},
		{ID: 2, URL: "ya.ru", Title: "yandex"},
	}
	req := httptest.NewRequest(http.MethodGet, "/api/v1/docs", nil)
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)
	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
	t.Log("Response: ", rr.Body)

	body := rr.Body.Bytes()
	var testDocs []index.Document

	if err := json.Unmarshal(body, &testDocs); err != nil {
		t.Fatal(string(body))
	}

	if !reflect.DeepEqual(testDocs, index.Documents) {
		t.Fatal(testDocs, index.Documents)
	}

}

func TestAPI_searchDoc(t *testing.T) {
	index.Documents = []index.Document{
		{ID: 1, URL: "url", Title: "title"},
		{ID: 2, URL: "ya.ru", Title: "yandex"},
	}
	index.MakeIndex()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/docs/yandex", nil)
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)
	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
	t.Log("Response: ", rr.Body)

	body := rr.Body.Bytes()
	var testTitle []string

	if err := json.Unmarshal(body, &testTitle); err != nil {
		t.Fatal(string(body))
	}

	if testTitle[0] != index.Documents[1].Title {
		t.Fatal(testTitle[0], index.Documents[1].Title)
	}
}

func TestAPI_deleteDoc(t *testing.T) {
	index.Documents = []index.Document{
		{ID: 1, URL: "url", Title: "title"},
		{ID: 2, URL: "ya.ru", Title: "yandex"},
	}
	lenDocs := len(index.Documents)
	title := index.Documents[lenDocs-1].Title
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/docs/1", nil)
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)

	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
	t.Log("Response: ", rr.Body)

	if lenDocs-1 != len(index.Documents) {
		t.Fatal(lenDocs-1, len(index.Documents))
	}

	if title != index.Documents[0].Title {
		t.Fatal(title, index.Documents[0].Title)
	}
}

func TestAPI_newDoc(t *testing.T) {
	index.Documents = []index.Document{
		{ID: 1, URL: "url", Title: "title"},
		{ID: 2, URL: "ya.ru", Title: "yandex"},
	}
	lenDocs := len(index.Documents)
	doc := index.Document{ID: 3, URL: "google.com", Title: "google"}
	payload, _ := json.Marshal(doc)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/docs", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)
	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	if lenDocs+1 != len(index.Documents) {
		t.Fatal(lenDocs+1, len(index.Documents))
	}

	if doc != index.Documents[lenDocs] {
		t.Fatal(doc, index.Documents[lenDocs])
	}
}

func TestAPI_updateDoc(t *testing.T) {
	index.Documents = []index.Document{
		{ID: 1, URL: "url", Title: "title"},
		{ID: 2, URL: "ya.ru", Title: "yandex"},
	}
	doc := index.Document{ID: 2, URL: "google.com", Title: "google"}
	payload, _ := json.Marshal(doc)
	req := httptest.NewRequest(http.MethodPut, "/api/v1/docs/2", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)
	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	if doc != index.Documents[len(index.Documents)-1] {
		t.Fatal(doc, index.Documents[len(index.Documents)-1])
	}
}
