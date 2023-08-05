package document

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	v1 "github.com/kinneko-de/api-contract/golang/kinnekode/restaurant/document/v1"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/httpheader"
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

	request, _ := http.NewRequest(http.MethodPost, expectedEndpoint, strings.NewReader(createValidRequest()))
	context.Request = request

	GeneratePreview(context)

	assert.EqualValues(t, http.StatusServiceUnavailable, response.Code)
}

func TestGeneratePreview_ValidRequest(t *testing.T) {
	mediaType := "application/pdf"
	size := uint64(134034)
	extension := ".pdf"
	expectedContentType := mediaType
	expectedContentLength := "134034";
	expectedContentDisposition := `attachment; filename="invoice.pdf"`;

	mockDocumentServiceGateway := mocks.NewDocumentServiceGateway(t)
	documentServiceGateway = mockDocumentServiceGateway
	mockClient := mocks.NewDocumentServiceClient(t)
	mockStream := mocks.NewDocumentService_GeneratePreviewClient(t)
	mockDocumentServiceGateway.SetupDocumentServiceGatewayToReturnClient(mockClient)
	mockClient.SetupGeneratePreview(mockStream)
	mockStream.EXPECT().Recv().Return(&v1.GeneratePreviewResponse{ File: &v1.GeneratePreviewResponse_Metadata{ Metadata: &v1.GeneratedFileMetadata{ MediaType: mediaType, Size: size, Extension: extension,}}}, nil).Once()
	mockStream.EXPECT().Recv().Return(&v1.GeneratePreviewResponse{ File: &v1.GeneratePreviewResponse_Chunk{ Chunk: make([]byte, 10)}}, nil).Once()
	mockStream.EXPECT().Recv().Return(nil, io.EOF).Once()
	mockStream.EXPECT().CloseSend().Return(nil).Once()

	response := httptest.NewRecorder()
	context := ginfixture.CreateContext(response);
	request, _ := http.NewRequest(http.MethodPost, expectedEndpoint, strings.NewReader(createValidRequest()))
	context.Request = request

	GeneratePreview(context)

	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, expectedContentType, response.Header().Get(httpheader.ContentType))
	assert.Equal(t, expectedContentLength, response.Header().Get(httpheader.ContentLength))
	assert.Equal(t, expectedContentDisposition, response.Header().Get(httpheader.ContentDisposition))
}

func createValidRequest() string {
	request := `{}`
	return request
}