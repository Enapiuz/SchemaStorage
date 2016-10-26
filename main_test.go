package main

import (
	"bytes"
	"encoding/json"
	"github.com/Enapiuz/SchemaStorage/core"
	"github.com/Enapiuz/SchemaStorage/handlers"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Schema(t *testing.T) {
	schema, err := ioutil.ReadFile("./test_suite/schema00.json")
	//document, err := ioutil.ReadFile("./test_suite/document00.json")

	server := core.InitializeCore()
	ts := httptest.NewServer(handlers.NewSchemaAddHandler(server))
	defer ts.Close()

	res, err := http.Post(ts.URL, "application/json", bytes.NewBuffer(schema))
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var dataObj core.Respomse
	err = json.Unmarshal(data, &dataObj)
	if err != nil {
		log.Fatal(err)
	}
	if dataObj.Ok == false {
		log.Fatal(dataObj.Info)
	}
}
