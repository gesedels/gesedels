// Package json implements JSON encoding and decoding functions.
package json

import (
	"encoding/json"
	"fmt"
)

// Decode returns a Go object from a decoded JSON string.
func Decode(text string, data any) error {
	if err := json.Unmarshal([]byte(text), data); err != nil {
		return fmt.Errorf("cannot decode %q - %w", text, err)
	}

	return nil
}

// DecodeMap returns a string map from a decoded JSON string.
func DecodeMap(text string) (map[string]string, error) {
	var smap = make(map[string]string)
	return smap, Decode(text, &smap)
}

// Encode returns an indented JSON string from an encoded object.
func Encode(data any) (string, error) {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", fmt.Errorf("cannot encode %#v - %w", data, err)
	}

	return string(bytes), nil
}
