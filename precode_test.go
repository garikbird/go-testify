package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req, _ := http.NewRequest("GET", "/?city=moscow&count=10", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	returnedCafes := strings.Split(responseRecorder.Body.String(), ",")
	actualCount := len(returnedCafes)

	assert.Equal(t, totalCount, actualCount)
}

func TestMainHandlerWhenCityNotSupported(t *testing.T) {
	req, _ := http.NewRequest("GET", "/?city=tarkov&count=2", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

func TestMainHandlerWhenRequestIsCorrect(t *testing.T) {
	req, _ := http.NewRequest("GET", "/?city=moscow&count=2", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String())
	assert.Equal(t, "Мир кофе,Сладкоежка", responseRecorder.Body.String())
}
