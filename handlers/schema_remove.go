package handlers

import (
	"github.com/Enapiuz/SchemaStorage/core"
	"github.com/Enapiuz/SchemaStorage/helpers/response"
	"github.com/Enapiuz/SchemaStorage/http_models"
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
			http_models.NewErrorResponseBody(http.StatusBadRequest, "Blank schema name"),
			http.StatusBadRequest)
		return
	}

	err := h.core.Repo.DeleteSchema(schemaName)
	if err != nil {
		response.Json(
			resp,
			http_models.NewErrorResponseBody(http.StatusNotFound, "Schema not found"),
			http.StatusNotFound)
		return
	}

	response.Json(resp, "", http.StatusOK)
}
