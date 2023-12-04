package router

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func SendRequestToSut(router *gin.Engine, req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}
