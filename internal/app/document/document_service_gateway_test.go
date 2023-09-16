package document

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadApiDocumentServiceConfig_HostIsMissing(t *testing.T) {
	t.Setenv(PortEnv, "8080")

	actualConfig, actualError := loadApiDocumentServiceConfig()

	assert.Equal(t, "", actualConfig)
	assert.NotNil(t, actualError)
	assert.Contains(t, actualError.Error(), HostEnv)
}

func TestLoadApiDocumentServiceConfig_PortIsMissing(t *testing.T) {
	t.Setenv(HostEnv, "http://localhost")

	actualConfig, actualError := loadApiDocumentServiceConfig()

	assert.Equal(t, "", actualConfig)
	assert.NotNil(t, actualError)
	assert.Contains(t, actualError.Error(), PortEnv)
}

func TestLoadApiDocumentServiceConfig_ValidConfig(t *testing.T) {
	t.Setenv(HostEnv, "http://localhost")
	t.Setenv(PortEnv, "8080")

	actualConfig, actualError := loadApiDocumentServiceConfig()

	assert.Equal(t, "http://localhost:8080", actualConfig)
	assert.Nil(t, actualError)
}

func TestCreateDocumentServiceClient_MissconfiguredUrl_ThrowsNoDialErrorMaybeBecauseItIsInsecure(t *testing.T) {
	t.Setenv(HostEnv, "iamnotthere")
	t.Setenv(PortEnv, "8080")
	ReadConfig()
	actualClient, actualError := documentServiceGateway.CreateDocumentServiceClient()

	assert.NotNil(t, actualClient)
	assert.Nil(t, actualError)
}
