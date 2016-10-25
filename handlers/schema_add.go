package handlers

import (
	"fmt"
	"github.com/Enapiuz/SchemaStorage/core"
	"github.com/Enapiuz/SchemaStorage/helpers/response"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type SchemaAddHandler struct {
	core *core.Core
}

func NewSchemaAddHandler(core *core.Core) *SchemaAddHandler {
	return &SchemaAddHandler{core: core}
}

func (h *SchemaAddHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	schemaName := vars["schemaname"]
	schemaString, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	if err != nil {
		response.Json(
			resp,
			core.Respomse{Ok: false, Info: "Can't read body"},
			http.StatusNotFound)
		return
	}

	if string(schemaString) == "" {
		response.Json(
			resp,
			core.Respomse{Ok: false, Info: "Blank request body"},
			http.StatusBadRequest)
		return
	}

	schemaData := core.NewSchema(schemaName, string(schemaString))

	collection := h.core.GetCollection(core.SchemaCollection)
	err = collection.Insert(schemaData)

	if err != nil {
		response.Json(
			resp,
			core.Respomse{Ok: false, Info: "Can't save schema"},
			http.StatusInternalServerError)
		return
	}

	response.Json(
		resp,
		core.Respomse{Ok: true, Info: fmt.Sprintf("Schema %s was added", schemaName)},
		http.StatusOK)
}
