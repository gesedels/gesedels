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
func AssertErr(t *testing.T, err error, regx string) {
	assert.Regexp(t, regx, err.Error())
}

// AssertResponse asserts the status code and JSON body from a ResponseRecorder.
func AssertResponse(t *testing.T, w *httptest.ResponseRecorder, code int, want any) {
	rslt := w.Result()
	body, err := io.ReadAll(rslt.Body)
	if err != nil {
		panic(err)
	}

	var data any
	err = json.Unmarshal(body, &data)
	assert.Equal(t, code, rslt.StatusCode)
	assert.Equal(t, want, data)
	assert.NoError(t, err)
}
