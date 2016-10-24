package handlers

import (
	"github.com/Enapiuz/SchemaStorage/helpers/response"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/Enapiuz/SchemaStorage/core"
)

type SchemaAddHandler struct {
	core *core.Core
}

func NewSchemaAddHandler(core *core.Core) *SchemaAddHandler {
	return &SchemaAddHandler{core: core}
}

func (h *SchemaAddHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	schemaname := vars["schemaname"]
	response.Json(
		resp,
		struct {
			Ok   bool
			Info string
		}{Ok: true, Info: fmt.Sprintf("Will add your schema `%s` to storage", schemaname)},
		http.StatusOK)
}
