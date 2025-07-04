///////////////////////////////////////////////////////////////////////////////////////
//            gesedels · a minimal key-value api in Go · by stephen malone           //
///////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/json"
	"net/http"
)

///////////////////////////////////////////////////////////////////////////////////////
//                          part one · constants and globals                         //
///////////////////////////////////////////////////////////////////////////////////////

// ResponseError is the default 500 Internal Error response.
const ResponseError = `{"status": "error", "message": "internal server error"}`

// ResponseType is the default HTTP response Content-Type.
const ResponseType = "application/json; charset=utf-8"

///////////////////////////////////////////////////////////////////////////////////////
//                         part two · http response functions                        //
///////////////////////////////////////////////////////////////////////////////////////

// writeJSON writes a marshalled JSON response to a ResponseWriter.
func writeJSON(w http.ResponseWriter, code int, data any) {
	body, err := json.Marshal(data)
	if err != nil {
		code = http.StatusInternalServerError
		body = []byte(ResponseError)
	}

	w.Header().Set("Content-Type", ResponseType)
	w.WriteHeader(code)
	w.Write(body)
}

// WriteError writes a JSend error response to a ResponseWriter.
func WriteError(w http.ResponseWriter, code int, data string) {
	writeJSON(w, code, map[string]any{"status": "error", "message": data})
}

// WriteFailure writes a JSend failure response to a ResponseWriter.
func WriteFailure(w http.ResponseWriter, code int, data map[string]string) {
	writeJSON(w, code, map[string]any{"status": "failure", "data": data})
}

// WriteSuccess writes a JSend success response to a ResponseWriter.
func WriteSuccess(w http.ResponseWriter, code int, data any) {
	writeJSON(w, code, map[string]any{"status": "success", "data": data})
}
