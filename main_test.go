package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое
func TestMainHendlerWhenStatusOkAndBodyNoNil(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// сервис возвращает код ответа 200
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// тело ответа не пустое
	assert.NotEmpty(t, responseRecorder.Body)
}

// Город, который передаётся в параметре city, не поддерживается
// Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа
func TestMainHandlerWhenNotCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=omsk", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Сервис возвращает код ответа 400
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	// wrong city value в теле ответа
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())

}

// Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	// Всего кафе в городе moscow
	totalCount := 4

	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	// Все доступные кафе
	assert.Len(t, list, totalCount)
}
