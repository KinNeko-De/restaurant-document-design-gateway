package testfixture

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func CreateGinContext(response *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(response)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	return ctx
}