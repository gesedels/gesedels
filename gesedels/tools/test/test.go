// Package test implements unit testing data and functions.
package test

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// AssertErr asserts an error message matches a regular expression.
func AssertErr(t *testing.T, err error, regx string) bool {
	return assert.Regexp(t, regx, err.Error())
}

// Response returns the status code and JSON body from a ResponseRecorder.
func Response(w *httptest.ResponseRecorder) (int, any) {
	rslt := w.Result()
	body, err := io.ReadAll(rslt.Body)
	if err != nil {
		panic(err)
	}

	var data any
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	return rslt.StatusCode, data
}
