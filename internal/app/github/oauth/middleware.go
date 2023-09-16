package oauth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	oauthgithub "golang.org/x/oauth2/github"
)

const ClientIdEnv = "GITHUBOAUTH_CLIENTID"
const ClientSecretEnv = "GITHUBOAUTH_CLIENTSECRET"

var (
	cache        sync.Map
	clientId     string = "set_by_init"
	clientSecret string = "set_by_init"
)

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
				log.Println(err)
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user can not be unauthorized. refresh the page without code and state"})
				return
			}
			ctx.Next()
		}
	}
}

func ReadConfig() (err error) {
	clientId, clientSecret, err = loadOAuthConfig()
	return err
}

func redirectToGithubOAuth(ctx *gin.Context) {
	githubOauthConfig := &oauth2.Config{
		RedirectURL:  getRedirectUrl(ctx),
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes:       []string{},
		Endpoint:     oauthgithub.Endpoint,
	}
	oauthStateString := strings.ReplaceAll(uuid.New().String(), "-", "")
	cache.Store(oauthStateString, time.Now())
	url := githubOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(ctx.Writer, ctx.Request, url, http.StatusTemporaryRedirect)
}

func getRedirectUrl(ctx *gin.Context) string {
	scheme := "http"
	if ctx.Request.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + path.Join(ctx.Request.Host, ctx.FullPath())
}

func writeUserIdToContext(ctx *gin.Context, state string, code string) error {
	userId, err := getUserId(ctx, state, code)
	if err != nil {
		return err
	}
	ctx.Set("userId", userId)
	return nil
}

func getUserId(ctx *gin.Context, state string, code string) (string, error) {
	value, loaded := cache.LoadAndDelete(state)
	if !loaded {
		return "", fmt.Errorf("state can not be loaded: %s", state)
	}
	if time.Since(value.(time.Time)) > 5*time.Minute {
		return "", fmt.Errorf("state is outdated: %s", state)
	}

	githubOauthConfig := &oauth2.Config{
		ClientID:     "555b51541211ad2920da",
		ClientSecret: "b8ea04669ded3c146e5c7e43aecb26ba133f7d32",
		Scopes:       []string{},
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

	contents, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		return "", fmt.Errorf("json.MarshlIndent() failed with %s", err.Error())
	}
	fmt.Printf("User:\n%s\n", string(contents)) // TODO: Debug
	fmt.Printf("User ID: %d", *user.ID)         // TODO: debug
	return strconv.FormatInt(*user.ID, 10), nil
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

func loadOAuthConfig() (clientId string, clientSecret string, err error) {
	clientId, found := os.LookupEnv(ClientIdEnv)
	if !found {
		err = errors.New("service host to generate documents is not configured. Expect environment variable " + ClientIdEnv)
		return
	}
	clientSecret, found = os.LookupEnv(ClientSecretEnv)
	if !found {
		err = errors.New("service port to generate documents is not configured. Expect environment variable " + ClientSecretEnv)
		return
	}

	return
}
