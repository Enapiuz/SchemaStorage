package handlers

import (
	"github.com/Enapiuz/SchemaStorage/core"
	"github.com/Enapiuz/SchemaStorage/helpers/response"
	"github.com/Enapiuz/SchemaStorage/http_models"
	"github.com/Enapiuz/SchemaStorage/validators"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type SchemaValidateHandler struct {
	core *core.Core
}

func NewSchemaValidateHandler(core *core.Core) *SchemaValidateHandler {
	return &SchemaValidateHandler{core: core}
}

func (h *SchemaValidateHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	schemaName := vars["schemaname"]
	if schemaName == "" {
		response.Json(
			resp,
			http_models.NewErrorResponseBody(http.StatusBadRequest, "Blank schema name"),
			http.StatusBadRequest)
		return
	}

	jsonBytes, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	if err != nil {
		response.Json(
			resp,
			http_models.NewErrorResponseBody(http.StatusBadRequest, "Blank request body"),
			http.StatusBadRequest)
		return
	}

	if len(jsonBytes) == 0 {
		response.Json(
			resp,
			http_models.NewErrorResponseBody(http.StatusBadRequest, "Blank request body"),
			http.StatusBadRequest)
		return
	}

	err = validators.ValidateJSONBytesBySchemaName(&jsonBytes, schemaName, h.core.Repo)
	if err != nil {
		response.Json(
			resp,
			http_models.NewErrorResponseBody(http.StatusBadRequest, err.Error()),
			http.StatusBadRequest)
	} else {
		response.Json(resp, "", http.StatusOK)
	}
}
