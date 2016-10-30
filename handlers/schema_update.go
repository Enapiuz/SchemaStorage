package handlers

import (
	"fmt"
	"github.com/Enapiuz/SchemaStorage/core"
	"github.com/Enapiuz/SchemaStorage/helpers/response"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strings"
)

type SchemaUpdateHandler struct {
	core *core.Core
}

func NewSchemaUpdateHandler(core *core.Core) *SchemaUpdateHandler {
	return &SchemaUpdateHandler{core: core}
}

func (h *SchemaUpdateHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	schemaName := strings.TrimSpace(vars["schemaname"])
	if schemaName == "" {
		response.Json(
			resp,
			core.Respomse{Ok: false, Info: "Blank schema name"},
			http.StatusBadRequest)
		return
	}

	schemaString, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	if err != nil {
		response.Json(
			resp,
			core.Respomse{Ok: false, Info: "Can't read body"},
			http.StatusBadRequest)
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
	err = collection.Update(struct {
		Name string
	}{Name: schemaName}, schemaData)

	if err != nil {
		response.Json(
			resp,
			core.Respomse{Ok: false, Info: "Can't update schema"},
			http.StatusInternalServerError)
		return
	}

	response.Json(
		resp,
		core.Respomse{Ok: true, Info: fmt.Sprintf("Schema '%s' was updated", schemaName)},
		http.StatusOK)
}
