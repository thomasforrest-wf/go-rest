/*
Copyright 2014 - 2015 Workiva, LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package rest

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Ensures that decodePayload returns an empty map for empty payloads.
func TestDecodePayloadEmpty(t *testing.T) {
	assert := assert.New(t)
	payload := bytes.NewBufferString("")

	decoded, err := decodePayload(payload.Bytes(), false)

	assert.Equal(Payload{}, decoded)
	assert.Nil(err)
}

// Ensures that decodePayload returns a nil and an error for invalid JSON payloads.
func TestDecodePayloadBadJSON(t *testing.T) {
	assert := assert.New(t)
	body := `{"foo": "bar", "baz": 1`
	payload := bytes.NewBufferString(body)

	decoded, err := decodePayload(payload.Bytes(), false)

	assert.Nil(decoded)
	assert.NotNil(err)
}

// Ensures that decodePayload returns a decoded map for JSON payloads.
func TestDecodePayloadHappyPath(t *testing.T) {
	assert := assert.New(t)
	body := `{"foo": "bar", "baz": 1}`
	payload := bytes.NewBufferString(body)

	decoded, err := decodePayload(payload.Bytes(), false)

	assert.Equal(Payload{"foo": "bar", "baz": float64(1)}, decoded)
	assert.Nil(err)
}

// Ensures that decodePayload returns a decoded map with json.Numbers for JSON payloads
func TestDecodePayloadHappyPathNumber(t *testing.T) {
	assert := assert.New(t)
	body := `{"foo": "bar", "baz": 1}`
	payload := bytes.NewBufferString(body)

	decoded, err := decodePayload(payload.Bytes(), true)

	var mockNumber json.Number

	mockNumber = "1"

	assert.Equal(Payload{"foo": "bar", "baz": mockNumber}, decoded)
	assert.Nil(err)
}

// Ensures that decodePayloadSlice returns an empty slice for empty payloads.
func TestDecodePayloadSliceEmpty(t *testing.T) {
	assert := assert.New(t)
	payload := bytes.NewBufferString("")

	decoded, err := decodePayloadSlice(payload.Bytes(), false)

	assert.Equal([]Payload{}, decoded)
	assert.Nil(err)
}

// Ensures that decodePayloadSlice returns a nil and an error for invalid JSON payloads.
func TestDecodePayloadSliceBadJSON(t *testing.T) {
	assert := assert.New(t)
	body := `[{"foo": "bar", "baz": 1`
	payload := bytes.NewBufferString(body)

	decoded, err := decodePayloadSlice(payload.Bytes(), false)

	assert.Nil(decoded)
	assert.NotNil(err)
}

// Ensures that decodePayloadSlice returns a decoded map for JSON payloads.
func TestDecodePayloadSliceHappyPath(t *testing.T) {
	assert := assert.New(t)
	body := `[{"foo": "bar", "baz": 1}]`
	payload := bytes.NewBufferString(body)

	decoded, err := decodePayloadSlice(payload.Bytes(), false)

	assert.Equal([]Payload{Payload{"foo": "bar", "baz": float64(1)}}, decoded)
	assert.Nil(err)
}

// Ensures that TestDecodePayloadSliceHappyPathNumber returns a decoded map with
// json.Numbers for JSON payloads
func TestDecodePayloadSliceHappyPathNumber(t *testing.T) {
	assert := assert.New(t)
	body := `[{"foo": "bar", "baz": 2}, {"bar": "foo", "baz": 3}]`
	payload := bytes.NewBufferString(body)

	decoded, err := decodePayloadSlice(payload.Bytes(), true)

	var mockNumber1 json.Number
	mockNumber1 = "2"
	var mockNumber2 json.Number
	mockNumber2 = "3"

	slicePayload := []Payload{
		Payload{"foo": "bar", "baz": mockNumber1},
		Payload{"bar": "foo", "baz": mockNumber2},
	}

	assert.Equal(slicePayload, decoded)
	assert.Nil(err)
}
