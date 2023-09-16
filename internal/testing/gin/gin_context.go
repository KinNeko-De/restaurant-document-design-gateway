package ginfixture

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func CreateContext(response *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(response)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	ctx.Set("userId", "1234567890")

	return ctx
}
