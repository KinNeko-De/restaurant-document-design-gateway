package document

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	v1 "github.com/kinneko-de/api-contract/golang/kinnekode/restaurant/document/v1"
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

func TestGeneratePreview_Valid(t *testing.T) {
	mockDocumentServiceGateway := mocks.NewDocumentServiceGateway(t)
	documentServiceGateway = mockDocumentServiceGateway
	mockClient := mocks.NewDocumentServiceClient(t)
	mockStream := mocks.NewDocumentService_GeneratePreviewClient(t)
	mockDocumentServiceGateway.SetupDocumentServiceGatewayToReturnClient(mockClient)
	mockClient.SetupGeneratePreview(mockStream)
	mockStream.EXPECT().Recv().Return(&v1.GeneratePreviewResponse{ File: &v1.GeneratePreviewResponse_Metadata{ Metadata: &v1.GeneratedFileMetadata{}}}, nil).Once()
	mockStream.EXPECT().Recv().Return(&v1.GeneratePreviewResponse{ File: &v1.GeneratePreviewResponse_Chunk{ Chunk: make([]byte, 10)}}, nil).Once()
	mockStream.EXPECT().Recv().Return(nil, io.EOF).Once()
	mockStream.EXPECT().CloseSend().Return(nil).Once()


	response := httptest.NewRecorder()
	context := ginfixture.CreateContext(response);
	request, _ := http.NewRequest(http.MethodPost, expectedEndpoint, strings.NewReader(createRequest()))
	context.Request = request

	GeneratePreview(context)

	assert.EqualValues(t, http.StatusCreated, response.Code)
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