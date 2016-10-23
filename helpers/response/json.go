package response

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	body, err := json.Marshal(data)
	if err != nil {
		body = []byte(`{"error": "Unknown error in encoding json!"}`)
	}

	w.Write(body)
}
