package test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.etcd.io/bbolt"
)

func TestAssertErr(t *testing.T) {
	// setup
	err := errors.New("test")

	// success
	AssertErr(t, err, "t.*")
}

func TestMockDB(t *testing.T) {
	// success
	db := MockDB(t)
	db.View(func(tx *bbolt.Tx) error {
		buck := tx.Bucket([]byte("foo"))
		assert.NotNil(t, buck)
		for item, want := range MockPair {
			data := buck.Get([]byte(item))
			assert.Equal(t, want, string(data))
		}

		return nil
	})
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
