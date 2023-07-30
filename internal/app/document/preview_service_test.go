package document

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kinneko-de/restaurant-document-design-gateway/internal/testing/ginfixture"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/testing/mocks"
	"github.com/stretchr/testify/assert"
)

const expectedEndpoint string = "/document/preview"

func TestGeneratePreview_RequestIsNil(t *testing.T) {
	response := httptest.NewRecorder()
	context := ginfixture.CreateContext(response);
	GeneratePreview(context)
	
	assert.EqualValues(t, http.StatusBadRequest, response.Code)
}

func TestGeneratePreview_DialError(t *testing.T) {
	mockDocumentServiceGateway := mocks.NewDocumentServiceGateway(t)
	mockDocumentServiceGateway.SetupDocumentServiceGatewayToReturnDialError()
	documentServiceGateway = mockDocumentServiceGateway

	response := httptest.NewRecorder()
	context := ginfixture.CreateContext(response);

	request, _ := http.NewRequest(http.MethodPost, expectedEndpoint, strings.NewReader(createRequest()))
	context.Request = request

	GeneratePreview(context)

	assert.EqualValues(t, http.StatusServiceUnavailable, response.Code)
}

/*
func TestValid(t *testing.T) {
	documentService := &mocks.DocumentServiceServer{}
	documentService.On("GeneratePreview", mock.AnythingOfType("v1.GeneratePreviewRequest"), mock.AnythingOfType("v1.DocumentService_GeneratePreviewServer")).Return(nil)
	
}
*/

func createRequest() string {
	request := `{}`
	return request
}