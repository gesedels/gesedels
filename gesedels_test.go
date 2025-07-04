///////////////////////////////////////////////////////////////////////////////////////
//               gesedels · unit tests and helpers · by stephen malone               //
///////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

///////////////////////////////////////////////////////////////////////////////////////
//                          part one · unit testing helpers                          //
///////////////////////////////////////////////////////////////////////////////////////

// assertError asserts an error message matches a regular expression.
func assertError(t *testing.T, err error, regx string) {
	assert.Regexp(t, regx, err.Error())
}

// assertResponse asserts the status code and JSON body from a ResponseRecorder.
func assertResponse(t *testing.T, w *httptest.ResponseRecorder, code int, want any) {
	var data any
	rslt := w.Result()
	body, _ := io.ReadAll(rslt.Body)
	err := json.Unmarshal(body, &data)
	assert.Equal(t, code, rslt.StatusCode)
	assert.Equal(t, want, data)
	assert.Equal(t, ResponseType, rslt.Header.Get("Content-Type"))
	assert.NoError(t, err)
}

///////////////////////////////////////////////////////////////////////////////////////
//                     part two · testing http response functions                    //
///////////////////////////////////////////////////////////////////////////////////////

func TestWriteJSON(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success - valid JSON
	writeJSON(w, http.StatusOK, "test")
	assertResponse(t, w, http.StatusOK, "test")

	// setup
	w = httptest.NewRecorder()

	// success - invalid JSON
	writeJSON(w, http.StatusOK, make(chan int))
	assertResponse(t, w, http.StatusInternalServerError, map[string]any{
		"status": "error", "message": "internal server error",
	})
}

func TestWriteError(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	WriteError(w, http.StatusInternalServerError, "test")
	assertResponse(t, w, http.StatusInternalServerError, map[string]any{
		"status": "error", "message": "test",
	})
}

func TestWriteFailure(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	WriteFailure(w, http.StatusBadRequest, map[string]string{"test": "test"})
	assertResponse(t, w, http.StatusBadRequest, map[string]any{
		"status": "failure", "data": map[string]any{"test": "test"},
	})
}

func TestWriteSuccess(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	WriteSuccess(w, http.StatusOK, "test")
	assertResponse(t, w, http.StatusOK, map[string]any{
		"status": "success", "data": "test",
	})
}
