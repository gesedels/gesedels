// Package test implements unit testing data and functions.
package test

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.etcd.io/bbolt"
)

// MockPair is a mock key-value pair for unit testing.
var MockPair = map[string]string{
	"body": "Bar.",
	"init": "1751435396",
	"hash": "8e2f3a93aeb2dff313fbb6e5b915261f36a8eca426fa7f8bd385f19c2ba287ae",
	"pass": "$2a$12$Bb1CsGvg7FP33U3XCse7tu5Z4VHP8sevkD7cKi8RQ.uyzGLYXxz76",
}

// AssertErr asserts an error message matches a regular expression.
func AssertErr(t *testing.T, err error, regx string) bool {
	return assert.Regexp(t, regx, err.Error())
}

// MockDB returns a temporary database populated with MockData.
func MockDB(t *testing.T) *bbolt.DB {
	path := filepath.Join(t.TempDir(), "bolt.db")
	db, _ := bbolt.Open(path, 0666, nil)
	db.Update(func(tx *bbolt.Tx) error {
		buck, _ := tx.CreateBucket([]byte("foo"))
		for item, data := range MockPair {
			buck.Put([]byte(item), []byte(data))
		}

		return nil
	})

	return db
}

// Response returns the status code and JSON body from a ResponseRecorder.
func Response(w *httptest.ResponseRecorder) (int, any) {
	var data any
	rslt := w.Result()
	body, _ := io.ReadAll(rslt.Body)
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	return rslt.StatusCode, data
}
