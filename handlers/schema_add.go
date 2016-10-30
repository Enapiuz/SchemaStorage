package handlers

import (
	"fmt"
	"github.com/Enapiuz/SchemaStorage/core"
	"github.com/Enapiuz/SchemaStorage/helpers/response"
	"github.com/Enapiuz/SchemaStorage/validators"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strings"
)

type SchemaAddHandler struct {
	core *core.Core
}

func NewSchemaAddHandler(core *core.Core) *SchemaAddHandler {
	return &SchemaAddHandler{core: core}
}

func (h *SchemaAddHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
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

	err = validators.ValidateBytes(&schemaString)
	if err != nil {
		response.Json(
			resp,
			core.Respomse{Ok: false, Info: "Invalid schema"},
			http.StatusInternalServerError)
		return
	}

	schemaData := core.NewSchema(schemaName, string(schemaString))

	collection := h.core.GetCollection(core.SchemaCollection)
	err = collection.Insert(schemaData)

	if err != nil {
		response.Json(
			resp,
			core.Respomse{Ok: false, Info: "Can't add schema, it's already exists"},
			http.StatusInternalServerError)
		return
	}

	response.Json(
		resp,
		core.Respomse{Ok: true, Info: fmt.Sprintf("Schema '%s' was added", schemaName)},
		http.StatusOK)
}
