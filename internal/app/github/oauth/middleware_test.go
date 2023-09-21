package oauth

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGithubOAuth_NoOAuthParameter_RedirectToGithub(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	ctx.Request = req

	GithubOAuth()(ctx)

	assert.Equal(t, http.StatusTemporaryRedirect, w.Code, "expected status Temporary Redirect; got %v", w.Code)
}

func TestGithubOAuthWithParameters(t *testing.T) {
	// Create a new Gin context for the test
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	// Call the middleware function with state and code parameters
	state := "teststate"
	code := "testcode"

	cache.Store(state, time.Now())
	cache.Store(code, time.Now())

	ctx.Request = httptest.NewRequest(http.MethodGet, "/?state="+state+"&code="+code, nil)
	GithubOAuth()(ctx)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("expected status OK; got %v", w.Code)
	}

	// Check the user ID in the context
	userId, exists := ctx.Get("userId")
	if !exists {
		t.Errorf("expected user ID in context; got none")
	}
	if userId != "123456" {
		t.Errorf("expected user ID 123456; got %v", userId)
	}
}
