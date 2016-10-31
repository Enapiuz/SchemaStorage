package handlers

import (
	"github.com/Enapiuz/SchemaStorage/core"
	"github.com/Enapiuz/SchemaStorage/db_models"
	"github.com/Enapiuz/SchemaStorage/helpers/response"
	"github.com/Enapiuz/SchemaStorage/http_models"
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
	defer req.Body.Close()
	vars := mux.Vars(req)
	schemaName := strings.TrimSpace(vars["schemaname"])
	if schemaName == "" {
		response.Json(
			resp,
			http_models.NewErrorResponseBody(http.StatusBadRequest, "Blank schema name"),
			http.StatusBadRequest)
		return
	}

	schemaBytes, err := ioutil.ReadAll(req.Body)

	if err != nil {
		response.Json(
			resp,
			http_models.NewErrorResponseBody(http.StatusBadRequest, err.Error()),
			http.StatusBadRequest)
		return
	}

	if len(schemaBytes) == 0 {
		response.Json(
			resp,
			http_models.NewErrorResponseBody(http.StatusBadRequest, "Blank request body"),
			http.StatusBadRequest)
		return
	}

	err = validators.ValidateJSONBytes(&schemaBytes)
	if err != nil {
		response.Json(
			resp,
			http_models.NewErrorResponseBody(http.StatusBadRequest, "Invalid schema"),
			http.StatusBadRequest)
		return
	}

	newSchema := db_models.NewSchema(schemaName, string(schemaBytes))
	err = h.core.Repo.InsertSchema(newSchema)
	if err != nil {
		response.Json(
			resp,
			http_models.NewErrorResponseBody(http.StatusConflict, "Schema already exists"),
			http.StatusConflict)
		return
	}

	response.Json(resp, "", http.StatusOK)
}
