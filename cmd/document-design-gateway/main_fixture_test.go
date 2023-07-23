package main

import (
	"net/http"
	"net/http/httptest"
)

func SendRequestToSut(req *http.Request) *httptest.ResponseRecorder {
	router := setupRouter()
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}