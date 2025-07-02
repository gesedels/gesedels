package json

import (
	"testing"

	"github.com/gesedels/gesedels/gesedels/tools/test"
	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	// setup
	var data int

	// success
	err := Decode("123", &data)
	assert.Equal(t, 123, data)
	assert.NoError(t, err)

	// error - cannot decode
	err = Decode("", &data)
	test.AssertErr(t, err, `cannot decode "" - unexpected end of JSON input`)
}

func TestDecodeMap(t *testing.T) {
	// success
	smap, err := DecodeMap(`{"foo": "bar"}`)
	assert.Equal(t, map[string]string{"foo": "bar"}, smap)
	assert.NoError(t, err)
}

func TestEncode(t *testing.T) {
	// success
	text, err := Encode([]int{123})
	assert.Equal(t, "[\n  123\n]", text)
	assert.NoError(t, err)

	// error - cannot encode
	text, err = Encode(make(chan int))
	assert.Empty(t, text)
	test.AssertErr(t, err, `cannot encode .* - json: unsupported type: chan int`)
}
