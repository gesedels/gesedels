package test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAssertErr(t *testing.T) {
	// setup
	err := errors.New("test error")

	// success
	AssertErr(t, err, "test .*")
}

func TestAssertResponse(t *testing.T) {
	// setup
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, `"test"`)

	// success
	AssertResponse(t, w, http.StatusOK, "test")
}
