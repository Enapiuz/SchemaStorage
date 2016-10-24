package handlers

import (
	"github.com/Enapiuz/SchemaStorage/helpers/response"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/Enapiuz/SchemaStorage/core"
)

type SchemaValidateHandler struct {
	core *core.Core
}

func NewSchemaValidateHandler(core *core.Core) *SchemaValidateHandler {
	return &SchemaValidateHandler{core: core}
}

func (h *SchemaValidateHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	schemaname := vars["schemaname"]
	response.Json(
		resp,
		struct {
			Ok   bool
			Info string
		}{Ok: true, Info: fmt.Sprintf("Will validate your json in body with schema `%s`", schemaname)},
		http.StatusOK)
}
