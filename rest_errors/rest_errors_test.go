package rest_errors

import (
	"encoding/json"
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
	"errors"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is the message", errors.New("database error"))
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
	assert.Equal(t, "this is the message", err.Message())
	assert.Equal(t, "message: this is the message - status: 500 - error: internal_server_error - causes: [database error]", err.Error())

	assert.NotNil(t, err.Causes)
	assert.Equal(t, 1, len(err.Causes()))
	assert.Equal(t, "database error", err.Causes()[0])
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("this is the message")
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status())
	assert.Equal(t, "this is the message", err.Message())
	assert.Equal(t, "message: this is the message - status: 400 - error: bad_request - causes: []", err.Error())

	assert.NotNil(t, err.Causes)
	assert.Equal(t, 0, len(err.Causes()))
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("this is the message")
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.Status())
	assert.Equal(t, "this is the message", err.Message())
	assert.Equal(t, "message: this is the message - status: 404 - error: not_found - causes: []", err.Error())

	assert.NotNil(t, err.Causes)
	assert.Equal(t, 0, len(err.Causes()))
}

func TestNewUnauthorizedError(t *testing.T) {
	err := NewUnauthorizedError("this is the message")
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusUnauthorized, err.Status())
	assert.Equal(t, "this is the message", err.Message())
	assert.Equal(t, "message: this is the message - status: 401 - error: unauthorized - causes: []", err.Error())

	assert.NotNil(t, err.Causes)
	assert.Equal(t, 0, len(err.Causes()))
}

func TestNewRestErrorFromBytes(t *testing.T) {
	err := NewUnauthorizedError("test_error")
	bytes, _ := json.Marshal(err)
	restErr, _ := NewRestErrorFromBytes(bytes)
	assert.NotNil(t, restErr)
	assert.Equal(t, http.StatusUnauthorized, restErr.Status())
	assert.Equal(t, "test_error", restErr.Message())
	assert.Equal(t, "message: test_error - status: 401 - error: unauthorized - causes: []", err.Error())

	assert.NotNil(t, restErr.Causes)
	assert.Equal(t, 0, len(restErr.Causes()))
}

func TestNewRestErrorFromBytesReturnsError(t *testing.T) {
	bytes := []byte{}
	restErr, err := NewRestErrorFromBytes(bytes)
	assert.Nil(t, restErr)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid json", err.Error())
}

func TestNewError(t *testing.T) {
	causes := []interface{}{errors.New("some_cause").Error()}
	err := NewRestError("this is the message", http.StatusBadGateway, "some_error", causes)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadGateway, err.Status())
	assert.Equal(t, "this is the message", err.Message())
	assert.Equal(t, "message: this is the message - status: 502 - error: some_error - causes: [some_cause]", err.Error())

	assert.NotNil(t, err.Causes)
	assert.Equal(t, 1, len(err.Causes()))
	assert.Equal(t, "some_cause", err.Causes()[0])
}