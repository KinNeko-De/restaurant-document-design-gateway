package main

import (
	"net/http"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/document"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/github/oauth"
	mainfixture "github.com/kinneko-de/restaurant-document-design-gateway/internal/testing/main"
)

const expectedEndpoint string = "/document/preview"

// tests that the service is mapped to the expected expectedEndpoint (no code coverage)
func TestDocumentPreview_RequestIsNil(t *testing.T) {
	t.Setenv(document.HostEnv, "http://localhost")
	t.Setenv(document.PortEnv, "8080")

	request, _ := http.NewRequest(http.MethodPost, expectedEndpoint, nil)

	response := mainfixture.SendRequestToSut(setupRouter(), request)

	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestDocumentPreview_GatewayConfigIsMissing(t *testing.T) {
	t.Skip("Does not run in github workflow")

	if os.Getenv("EXECUTE") == "1" {
		main()
		return
	}

	t.Setenv(oauth.ClientIdEnv, "1234567890")
	t.Setenv(oauth.ClientSecretEnv, "1234567890")
	cmd := exec.Command(os.Args[0], "-test.run=TestDocumentPreview_GatewayConfigIsMissing")
	cmd.Env = append(os.Environ(), "EXECUTE=1")
	err := cmd.Run()
	require.NotNil(t, err)
	exitCode := err.(*exec.ExitError).ExitCode()
	assert.Equal(t, 1, exitCode)
}

func TestDocumentPreview_OAuthConfigIsMissing(t *testing.T) {
	t.Skip("Does not run in github workflow")

	if os.Getenv("EXECUTE") == "1" {
		main()
		return
	}

	t.Setenv(document.HostEnv, "http://localhost")
	t.Setenv(document.PortEnv, "8080")
	cmd := exec.Command(os.Args[0], "-test.run=TestDocumentPreview_OAuthConfigIsMissing")
	cmd.Env = append(os.Environ(), "EXECUTE=1")
	err := cmd.Run()
	require.NotNil(t, err)
	exitCode := err.(*exec.ExitError).ExitCode()
	assert.Equal(t, 1, exitCode)
}
