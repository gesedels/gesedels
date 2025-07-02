package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gesedels/gesedels/gesedels/tools/test"
	"github.com/stretchr/testify/assert"
)

func TestWriteJSON(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	writeJSON(w, http.StatusOK, "test")
	code, data := test.Response(w)
	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, "test", data)
}

func TestError(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	Error(w, http.StatusInternalServerError, "test")
	code, data := test.Response(w)
	assert.Equal(t, http.StatusInternalServerError, code)
	assert.Equal(t, map[string]any{
		"status":  "error",
		"message": "test",
	}, data)
}

func TestFailure(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	Failure(w, http.StatusBadRequest, map[string]string{"test": "test"})
	code, data := test.Response(w)
	assert.Equal(t, http.StatusBadRequest, code)
	assert.Equal(t, map[string]any{
		"status": "failure",
		"data":   map[string]any{"test": "test"},
	}, data)
}

func TestSuccess(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	Success(w, http.StatusOK, "test")
	code, data := test.Response(w)
	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, map[string]any{
		"status": "success",
		"data":   "test",
	}, data)
}
