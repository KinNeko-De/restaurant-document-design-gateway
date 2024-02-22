package oauth

import (
	"context"
	"fmt"
	"net/http"
	"path"
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

const emailScope = "user:email"
const UserContextKey = "userId"

func GithubOAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		state := ctx.Request.FormValue("state")
		code := ctx.Request.FormValue("code")
		if state == "" || code == "" {
			redirectToGithubOAuth(ctx)
			return
		} else {
			err := writeUserIdToContext(ctx, state, code)
			if err != nil {
				logger.Logger.Error().Err(err).Msg("Failed to write user id to context")
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user can not be unauthorized. refresh the page without code and state"})
				return
			}
			ctx.Next()
		}
	}
}

func redirectToGithubOAuth(ctx *gin.Context) {
	githubOauthConfig := &oauth2.Config{
		RedirectURL:  getRedirectUrl(ctx),
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes:       []string{emailScope},
		Endpoint:     oauthgithub.Endpoint,
	}
	oauthStateString := strings.ReplaceAll(uuid.New().String(), "-", "")
	cache.Store(oauthStateString, time.Now())
	url := githubOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(ctx.Writer, ctx.Request, url, http.StatusTemporaryRedirect)
	ctx.Abort()
}

func getRedirectUrl(ctx *gin.Context) string {
	scheme := "http"
	if ctx.Request.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + path.Join(ctx.Request.Host, ctx.FullPath())
}

func writeUserIdToContext(ctx *gin.Context, state string, code string) error {
	userEmail, err := getUserEmail(ctx, state, code)
	if err != nil {
		return err
	}
	ctx.Set(UserContextKey, userEmail)
	return nil
}

func getUserEmail(ctx *gin.Context, state string, code string) (string, error) {
	value, loaded := cache.LoadAndDelete(state)
	if !loaded {
		return "", fmt.Errorf("state can not be loaded: %s", state)
	}
	if time.Since(value.(time.Time)) > 5*time.Minute {
		return "", fmt.Errorf("state is outdated: %s", state)
	}

	githubOauthConfig := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes:       []string{emailScope},
		Endpoint:     oauthgithub.Endpoint,
	}
	token, err := githubOauthConfig.Exchange(context.Background(), code)
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

	logger.Logger.Debug().Msgf("User Email: %s", user.GetEmail())
	return user.GetEmail(), nil
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
