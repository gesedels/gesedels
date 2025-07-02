package test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssertErr(t *testing.T) {
	// setup
	err := errors.New("test error")

	// success
	AssertErr(t, err, "test .*")
}

func TestResponse(t *testing.T) {
	// setup
	w := httptest.NewRecorder()
	fmt.Fprintf(w, `"test"`)

	// success
	code, data := Response(w)
	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, "test", data)
}
