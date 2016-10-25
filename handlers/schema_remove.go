package handlers

import (
	"fmt"
	"github.com/Enapiuz/SchemaStorage/core"
	"github.com/Enapiuz/SchemaStorage/helpers/response"
	"github.com/gorilla/mux"
	"net/http"
)

type SchemaRemoveHandler struct {
	core *core.Core
}

func NewSchemaRemoveHandler(core *core.Core) *SchemaRemoveHandler {
	return &SchemaRemoveHandler{core: core}
}

func (h *SchemaRemoveHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	schemaName := vars["schemaname"]
	if schemaName == "" {
		response.Json(
			resp,
			core.Respomse{Ok: false, Info: "Blank schema name"},
			http.StatusBadRequest)
		return
	}

	err := h.core.GetCollection(core.SchemaCollection).Remove(struct {
		Name string
	}{Name: schemaName})

	if err != nil {
		response.Json(
			resp,
			core.Respomse{Ok: false, Info: fmt.Sprintf("Can't delete schema '%s'", schemaName)},
			http.StatusNotFound)
		return
	}

	response.Json(
		resp,
		core.Respomse{Ok: true, Info: fmt.Sprintf("Schema '%s' was deleted", schemaName)},
		http.StatusOK)
}
