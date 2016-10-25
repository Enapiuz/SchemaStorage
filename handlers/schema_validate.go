package handlers

import (
	"fmt"
	"github.com/Enapiuz/SchemaStorage/core"
	"github.com/Enapiuz/SchemaStorage/helpers/response"
	"github.com/gorilla/mux"
	"github.com/xeipuuv/gojsonschema"
	"io/ioutil"
	"net/http"
)

type SchemaValidateHandler struct {
	core *core.Core
}

func NewSchemaValidateHandler(core *core.Core) *SchemaValidateHandler {
	return &SchemaValidateHandler{core: core}
}

func (h *SchemaValidateHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	schemaName := vars["schemaname"]
	if schemaName == "" {
		response.Json(
			resp,
			core.Respomse{Ok: false, Info: "Blank schema name"},
			http.StatusBadRequest)
		return
	}

	jsonString, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	if err != nil {
		response.Json(
			resp,
			core.Respomse{Ok: false, Info: "Can't read body"},
			http.StatusBadRequest)
		return
	}

	if string(jsonString) == "" {
		response.Json(
			resp,
			core.Respomse{Ok: false, Info: "Blank request body"},
			http.StatusBadRequest)
		return
	}

	var schema core.Schema
	err = h.core.GetCollection(core.SchemaCollection).Find(struct {
		Name string
	}{Name: schemaName}).One(&schema)
	if err != nil {
		response.Json(
			resp,
			core.Respomse{Ok: false, Info: fmt.Sprintf("Can't find schema '%s'", schemaName)},
			http.StatusNotFound)
		return
	}

	schemaLoader := gojsonschema.NewStringLoader(schema.Data)
	documentLoader := gojsonschema.NewBytesLoader(jsonString)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)

	if err != nil {
		response.Json(
			resp,
			core.Respomse{Ok: false, Info: fmt.Sprintf("Unknown error '%s'", err.Error())},
			http.StatusNotFound)
		return
	}

	if result.Valid() {
		response.Json(
			resp,
			core.Respomse{Ok: true, Info: "Document is valid"},
			http.StatusOK)
		return
	} else {
		//fmt.Printf("The document is not valid. see errors :\n")
		//for _, desc := range result.Errors() {
		//	fmt.Printf("- %s\n", desc)
		//}
		response.Json(
			resp,
			core.Respomse{Ok: true, Info: "Document is not valid"},
			http.StatusOK)
		return
	}
}
