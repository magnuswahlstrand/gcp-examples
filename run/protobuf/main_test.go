package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello", nil)
	rr := httptest.NewRecorder()

	os.Setenv("NAME", "Space")

	// Act
	helloHandler(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	assert.Equal(t, "Hello Space!\n", rr.Body.String())
}

func TestHandlerDefault(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello", nil)
	rr := httptest.NewRecorder()

	os.Setenv("NAME", "")

	// Act
	helloHandler(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	assert.Equal(t, "Hello 世界!\n", rr.Body.String())
}

func TestHandlerPanic(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello", nil)
	rr := httptest.NewRecorder()

	os.Setenv("NAME", "Mars")

	// Assert
	defer func() {
		if r := recover(); r == nil {
			require.Fail(t, "expected panic, got no panic")
		}
	}()

	// Act
	helloHandler(rr, req)
}
