package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gesedels/gesedels/gesedels/tools/test"
)

func TestWriteJSON(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success - valid JSON
	writeJSON(w, http.StatusOK, "test")
	test.AssertResponse(t, w, http.StatusOK, "test")

	// setup
	w = httptest.NewRecorder()

	// success - invalid JSON
	writeJSON(w, http.StatusOK, make(chan int))
	test.AssertResponse(t, w, http.StatusInternalServerError, map[string]any{
		"status":  "error",
		"message": "internal server error",
	})
}

func TestError(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	Error(w, http.StatusInternalServerError, "test")
	test.AssertResponse(t, w, http.StatusInternalServerError, map[string]any{
		"status":  "error",
		"message": "test",
	})
}

func TestFailure(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	Failure(w, http.StatusBadRequest, map[string]string{"test": "test"})
	test.AssertResponse(t, w, http.StatusBadRequest, map[string]any{
		"status": "failure",
		"data":   map[string]any{"test": "test"},
	})
}

func TestSuccess(t *testing.T) {
	// setup
	w := httptest.NewRecorder()

	// success
	Success(w, http.StatusOK, "test")
	test.AssertResponse(t, w, http.StatusOK, map[string]any{
		"status": "success",
		"data":   "test",
	})
}
