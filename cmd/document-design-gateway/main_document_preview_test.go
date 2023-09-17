package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/document"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/github/oauth"
	mainfixture "github.com/kinneko-de/restaurant-document-design-gateway/internal/testing/main"
)

// tests that the service is mapped to the expected expectedEndpoint (no code coverage)
func TestDocumentPreview_DocumentPreviewIsMapped(t *testing.T) {
	const expectedEndpoint string = "/document/preview"

	t.Setenv(document.HostEnv, "http://localhost")
	t.Setenv(document.PortEnv, "8080")
	t.Setenv(oauth.ClientIdEnv, "1234567890")
	t.Setenv(oauth.ClientSecretEnv, "1234567890")

	request, _ := http.NewRequest(http.MethodPost, expectedEndpoint, nil)

	response := mainfixture.SendRequestToSut(setupRouter(), request)

	// because of authentication, the response should be a redirect
	assert.Equal(t, http.StatusTemporaryRedirect, response.Code)
}
