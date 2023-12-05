package oauth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
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
	req := httptest.NewRequest(http.MethodGet, requestedUri, nil)
	ctx.Request = req

	GithubOAuth(CreateOAuthFunc(ctx))(ctx)

	assert.Equal(t, http.StatusTemporaryRedirect, w.Code)
	url := w.Header().Get("Location")

	assert.Regexp(t, regexp.MustCompile(expectedRedirectUri), url)
}

func TestGithubOAuth_WithValidParameters(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	state := "teststate"
	code := "testcode"
	cache.Store(state, time.Now())
	cache.Store(code, time.Now())

	ctx.Request = httptest.NewRequest(http.MethodGet, "/?state="+state+"&code="+code, nil)
	GithubOAuth(&OAuth2Mock{})(ctx)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the user ID in the context
	userId, exists := ctx.Get("userId")
	if !exists {
		t.Errorf("expected user ID in context; got none")
	}
	if userId != "123456" {
		t.Errorf("expected user ID 123456; got %v", userId)
	}
}

type OAuth2Mock struct{}

// AuthCodeURL redirects to our own server.
func (o *OAuth2Mock) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	u := url.URL{
		Scheme: "http",
		Host:   "localhost",
		Path:   "login/oauth/authorize",
	}

	v := url.Values{}
	v.Set("state", state)

	u.RawQuery = v.Encode()
	return u.String()
}

// Exchange takes the code and returns a real token.
func (o *OAuth2Mock) Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return &oauth2.Token{
		AccessToken: "AccessToken",
		Expiry:      time.Now().Add(1 * time.Hour),
	}, nil
}

// Client returns a new http.Client.
func (o *OAuth2Mock) Client(ctx context.Context, t *oauth2.Token) *http.Client {
	return &http.Client{}
}
