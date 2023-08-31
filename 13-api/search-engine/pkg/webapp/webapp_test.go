package webapp

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"

	"go-core-4/11-net/search-engine/pkg/index"
)

var testMux *mux.Router

func TestMain(m *testing.M) {
	index.Documents = []index.Document{{ID: 0, URL: "url", Title: "title"}}
	index.MakeIndex()
	testMux = mux.NewRouter()
	endpoints(testMux)
	m.Run()
}

func Test_mainHandler_docs(t *testing.T) {
	data := []int{}
	payload, _ := json.Marshal(data)

	req := httptest.NewRequest(http.MethodPost, "/docs", bytes.NewBuffer(payload))

	rr := httptest.NewRecorder()

	testMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	t.Log("Response: ", rr.Body)
	//=========================================================

	req = httptest.NewRequest(http.MethodGet, "/docs", nil)
	rr = httptest.NewRecorder()

	testMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	body := rr.Body.Bytes()
	docs := make([]index.Document, 1)

	if err := json.Unmarshal(body, &docs); err != nil {
		t.Fatal(string(body))
	}

	if !reflect.DeepEqual(docs, index.Documents) {
		t.Fatal(docs, index.Documents)
	}
}

func Test_mainHandler_index(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/index", nil)
	rr := httptest.NewRecorder()

	testMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	body := rr.Body.Bytes()
	var testIndex index.Index

	if err := json.Unmarshal(body, &testIndex); err != nil {
		t.Fatal(string(body))
	}

	if !reflect.DeepEqual(testIndex, index.GetIndex()) {
		t.Fatal(testIndex, index.GetIndex())
	}
}
