package oauth

import (
	"context"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
	"github.com/google/uuid"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/operation/logger"
	"golang.org/x/oauth2"
	oauthgithub "golang.org/x/oauth2/github"
)

var (
	cache sync.Map
)

func GithubOAuth(config OAuth2ConfigInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request == nil {
			redirectToGithubOAuth(ctx, config)
			return
		}

		state := ctx.Request.FormValue("state")
		code := ctx.Request.FormValue("code")
		if state == "" || code == "" {
			redirectToGithubOAuth(ctx, config)
			return
		} else {
			err := writeUserIdToContext(ctx, config, state, code)
			if err != nil {
				logger.Logger.Error().Err(err).Msg("Failed to write user id to context")
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user can not be unauthorized. refresh the page without code and state"})
				return
			}
			ctx.Next()
		}
	}
}

func redirectToGithubOAuth(ctx *gin.Context, config OAuth2ConfigInterface) {
	oauthStateString := strings.ReplaceAll(uuid.New().String(), "-", "")
	cache.Store(oauthStateString, time.Now())
	url := config.AuthCodeURL(oauthStateString)
	http.Redirect(ctx.Writer, ctx.Request, url, http.StatusTemporaryRedirect)
	ctx.Abort()
}

func getRedirectUrl(ctx *gin.Context) string {
	scheme := "http"
	if ctx.Request.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + path.Join(ctx.Request.Host, ctx.Request.URL.Path)
}

func writeUserIdToContext(ctx *gin.Context, config OAuth2ConfigInterface, state string, code string) error {
	userId, err := getUserId(ctx, config, state, code)
	if err != nil {
		return err
	}

	ctx.Set("userId", userId)
	return nil
}

func getUserId(ctx *gin.Context, config OAuth2ConfigInterface, state string, code string) (string, error) {
	value, loaded := cache.LoadAndDelete(state)
	if !loaded {
		return "", fmt.Errorf("state can not be loaded: %s", state)
	}
	if time.Since(value.(time.Time)) > 5*time.Minute {
		return "", fmt.Errorf("state is outdated: %s", state)
	}

	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		return "", fmt.Errorf("code exchange failed: %s", err.Error())
	}

	tokenSource := &TokenSource{
		AccessToken: token.AccessToken,
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := github.NewClient(oauthClient)

	user, _, err := client.Users.Get(ctx, "")
	if err != nil {
		return "", fmt.Errorf("client.Users.Get() faled with '%s'", err.Error())
	}

	logger.Logger.Debug().Msgf("User ID: %d, User: %s", *user.ID, user.String())
	return strconv.FormatInt(*user.ID, 10), nil
}

func CreateOAuthFunc(ctx *gin.Context) *oauth2.Config {
	githubOauthConfig := &oauth2.Config{
		RedirectURL:  getRedirectUrl(ctx),
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes:       []string{},
		Endpoint:     oauthgithub.Endpoint,
	}
	return githubOauthConfig
}

type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

type OAuth2ConfigInterface interface {
	AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string
	Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
	Client(ctx context.Context, t *oauth2.Token) *http.Client
}
