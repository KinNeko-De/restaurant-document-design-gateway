package document

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadApiDocumentServiceConfig_HostIsMissing(t *testing.T) {
	t.Setenv(portEnv, "8080")

	actualConfig, actualError := loadApiDocumentServiceConfig()

	assert.Equal(t, "", actualConfig)
	assert.NotNil(t, actualError)
	assert.Contains(t, actualError.Error(), hostEnv)
}

func TestLoadApiDocumentServiceConfig_PortIsMissing(t *testing.T) {
	t.Setenv(hostEnv, "http://localhost")

	actualConfig, actualError := loadApiDocumentServiceConfig()

	assert.Equal(t, "", actualConfig)
	assert.NotNil(t, actualError)
	assert.Contains(t, actualError.Error(), portEnv)
}

func TestLoadApiDocumentServiceConfig_ValidConfig(t *testing.T) {
	t.Setenv(hostEnv, "http://localhost")
	t.Setenv(portEnv, "8080")

	actualConfig, actualError := loadApiDocumentServiceConfig()

	assert.Equal(t, "http://localhost:8080", actualConfig)
	assert.Nil(t, actualError)
}

func TestCreateDocumentServiceClient_MissconfiguredUrl_ThrowsNoDialErrorMaybeBecauseItIsInsecure(t *testing.T) {
	t.Setenv(hostEnv, "iamnotthere")
	t.Setenv(portEnv, "8080")
	ReadConfig()
	actualClient, actualError := documentServiceGateway.CreateDocumentServiceClient()

	assert.NotNil(t, actualClient)
	assert.Nil(t, actualError)
}
