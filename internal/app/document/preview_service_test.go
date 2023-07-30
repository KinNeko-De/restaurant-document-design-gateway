package document

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kinneko-de/restaurant-document-design-gateway/internal/testing/ginfixture"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/testing/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePreview_RequestIsNil(t *testing.T) {
	response := httptest.NewRecorder()
	context := ginfixture.CreateContext(response);
	GeneratePreview(context)
	
	assert.EqualValues(t, http.StatusBadRequest, response.Code)
}


func TestGeneratePreview_DialError(t *testing.T) {
	mockDocumentServiceGateway := &mocks.DocumentServiceGateway{}
	mockDocumentServiceGateway.On("CreateDocumentServiceClient").Return(nil, errors.New("I want to see this error!"))
	documentServiceGateway = mockDocumentServiceGateway

	response := httptest.NewRecorder()
	context := ginfixture.CreateContext(response);

	requestJson := createRequest()
	request, _ := http.NewRequest(http.MethodPost, "/document/preview", strings.NewReader(requestJson))
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