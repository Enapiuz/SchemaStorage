package handlers

import (
	"github.com/Enapiuz/SchemaStorage/helpers/response"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

type SchemaAddHandler struct {
}

func NewSchemaAddHandler() *SchemaAddHandler {
	return &SchemaAddHandler{}
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
