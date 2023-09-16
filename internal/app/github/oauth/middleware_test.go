package oauth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig_EverythingIsMissing(t *testing.T) {
	actualError := ReadConfig()
	assert.NotNil(t, actualError)
}

func TestLoadApiDocumentServiceConfig_ClientIdIsMissing(t *testing.T) {
	t.Setenv(ClientSecretEnv, "12345")

	actualCLientId, actualClientSecret, actualError := loadOAuthConfig()

	assert.Equal(t, "", actualCLientId)
	assert.Equal(t, "", actualClientSecret)
	assert.NotNil(t, actualError)
	assert.Contains(t, actualError.Error(), ClientIdEnv)
}

func TestLoadApiDocumentServiceConfig_PortIsMissing(t *testing.T) {
	t.Setenv(ClientIdEnv, "12345675")

	actualCLientId, actualClientSecret, actualError := loadOAuthConfig()

	assert.Equal(t, "12345675", actualCLientId)
	assert.Equal(t, "", actualClientSecret)
	assert.NotNil(t, actualError)
	assert.Contains(t, actualError.Error(), ClientSecretEnv)
}

func TestLoadApiDocumentServiceConfig_ValidConfig(t *testing.T) {
	t.Setenv(ClientIdEnv, "12345675")
	t.Setenv(ClientSecretEnv, "12345")

	actualCLientId, actualClientSecret, actualError := loadOAuthConfig()

	assert.Equal(t, "12345675", actualCLientId)
	assert.Equal(t, "12345", actualClientSecret)
	assert.Nil(t, actualError)
}
