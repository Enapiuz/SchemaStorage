package handlers

import (
	"github.com/Enapiuz/SchemaStorage/helpers/response"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/Enapiuz/SchemaStorage/core"
)

type SchemaRemoveHandler struct {
	core *core.Core
}

func NewSchemaRemoveHandler(core *core.Core) *SchemaRemoveHandler {
	return &SchemaRemoveHandler{core: core}
}

func (h *SchemaRemoveHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	schemaname := vars["schemaname"]
	response.Json(
		resp,
		struct {
			Ok   bool
			Info string
		}{Ok: true, Info: fmt.Sprintf("Will remove your schema `%s` from database", schemaname)},
		http.StatusOK)
}
