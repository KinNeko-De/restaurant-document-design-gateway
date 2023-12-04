package oauth

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGithubOAuth_NoOAuthParameter_RedirectToGithub(t *testing.T) {
	expectedClientId := "12345675"
	t.Setenv(ClientIdEnv, expectedClientId)
	t.Setenv(ClientSecretEnv, "12345")
	ReadConfig()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	requestedUri := "http://localhost:8080/document/preview/demo"
	expectedRequestedUri := url.QueryEscape(requestedUri)
	expectedGithubOAuthUri := "https://github.com/login/oauth/authorize?"
	expectedRedirectUri := regexp.QuoteMeta(expectedGithubOAuthUri) + "client_id=" + expectedClientId + "&redirect_uri=" + expectedRequestedUri + "&response_type=code&state=[a-z0-9]{32}"
	req, _ := http.NewRequest(http.MethodGet, requestedUri, nil)
	ctx.Request = req

	GithubOAuth()(ctx)

	assert.Equal(t, http.StatusTemporaryRedirect, w.Code)
	url := w.Header().Get("Location")

	assert.Regexp(t, regexp.MustCompile(expectedRedirectUri), url)
}

/*
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
*/
