package http_models

type ErrorResponseBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewErrorResponseBody(code int, message string) *ErrorResponseBody {
	return &ErrorResponseBody{Code: code, Message: message}
}
