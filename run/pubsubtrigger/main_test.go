package main

import (
	"encoding/base64"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)



func TestHandler(t *testing.T) {
	data := base64.StdEncoding.EncodeToString([]byte("Badger"))
	payload := fmt.Sprintf(`{"message":{"data":"%s","id":"some-id"}}`, data)

	req := httptest.NewRequest(http.MethodPost, "/subscribe", strings.NewReader(payload))
	rr := httptest.NewRecorder()

	// Act
	subscribePubSubHandler(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	assert.Equal(t, "some-id: Badger\n", rr.Body.String())
}

func TestHandlerErrorNotBase64(t *testing.T) {
	data := "Badger"
	payload := fmt.Sprintf(`{"message":{"data":"%s","id":"some-id"}}`, data)

	req := httptest.NewRequest(http.MethodPost, "/subscribe", strings.NewReader(payload))
	rr := httptest.NewRecorder()

	// Act
	subscribePubSubHandler(rr, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, rr.Result().StatusCode)
}