// Package http implements HTTP response functions.
package http

import (
	"fmt"
	"net/http"

	"github.com/gesedels/gesedels/gesedels/tools/json"
)

// writeJSON writes a JSON HTTP response to a ResponseWriter.
func writeJSON(w http.ResponseWriter, code int, data any) {
	text, err := json.Encode(data)
	if err != nil {
		code = http.StatusInternalServerError
		text = `{"status": "error", "message": "internal server error"}`
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprint(w, text)
}

// Error writes a JSend error response to a ResponseWriter.
func Error(w http.ResponseWriter, code int, data string) {
	writeJSON(w, code, map[string]any{
		"status":  "error",
		"message": data,
	})
}

// Failure writes a JSend failure response to a ResponseWriter.
func Failure(w http.ResponseWriter, code int, data map[string]string) {
	writeJSON(w, code, map[string]any{
		"status": "failure",
		"data":   data,
	})
}

// Success writes a JSend success response to a ResponseWriter.
func Success(w http.ResponseWriter, code int, data any) {
	writeJSON(w, code, map[string]any{
		"status": "success",
		"data":   data,
	})
}
