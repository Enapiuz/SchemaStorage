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

type SchemaUpdateHandler struct {
	core *core.Core
}

func NewSchemaUpdateHandler(core *core.Core) *SchemaUpdateHandler {
	return &SchemaUpdateHandler{core: core}
}

func (h *SchemaUpdateHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
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

	schemaData := db_models.NewSchema(schemaName, string(schemaBytes))
	err = h.core.Repo.UpdateSchema(schemaName, schemaData)
	if err != nil {
		var statusCode int
		if err.Error() == "not found" {
			statusCode = http.StatusNotFound
		} else {
			statusCode = http.StatusBadRequest
		}
		response.Json(resp, http_models.NewErrorResponseBody(statusCode, err.Error()), statusCode)
		return
	}

	response.Json(resp, "", http.StatusOK)
}
