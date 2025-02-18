// response.go
package models

// ResponseMessage for success
type ResponseMessage struct {
	Message string `json:"message"`
}

// ErrorResponse for error messages
type ErrorResponse struct {
	Error string `json:"error"`
}
