package oauth

import (
	"errors"
	"os"
)

const ClientIdEnv = "GITHUBOAUTH_CLIENTID"
const ClientSecretEnv = "GITHUBOAUTH_CLIENTSECRET"

var (
	clientId     string = "set_by_init"
	clientSecret string = "set_by_init"
)

func ReadConfig() (err error) {
	clientId, clientSecret, err = loadOAuthConfig()
	return err
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
