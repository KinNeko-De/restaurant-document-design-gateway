package document

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kinneko-de/restaurant-document-design-gateway/internal/httpheader"
	fixture "github.com/kinneko-de/restaurant-document-design-gateway/internal/testing/document"
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
	expectedFile := []byte{84,104,101,32,97,110,115,119,101,114,32,105,115,32,52,50}

	mockDocumentServiceGateway := mocks.NewDocumentServiceGateway(t)
	documentServiceGateway = mockDocumentServiceGateway
	mockClient := mocks.NewDocumentServiceClient(t)
	mockStream := mocks.NewDocumentService_GeneratePreviewClient(t)
	mockDocumentServiceGateway.SetupDocumentServiceGatewayToReturnClient(mockClient)
	mockClient.SetupGeneratePreview(mockStream)
	mockStream.EXPECT().Recv().
	Return(fixture.NewGeneratePreviewResponseMetadataBuilder().
		WithMediaType(mediaType).
		WithSize(size).
		WithExtension(extension).
		Build(), nil).Once()
	mockStream.EXPECT().Recv().Return(fixture.NewGeneratePreviewResponseChunkBuilder().WithChunk(expectedFile[0:6]).Build(), nil).Once()
	mockStream.EXPECT().Recv().Return(fixture.NewGeneratePreviewResponseChunkBuilder().WithChunk(expectedFile[6:11]).Build(), nil).Once()
	mockStream.EXPECT().Recv().Return(fixture.NewGeneratePreviewResponseChunkBuilder().WithChunk(expectedFile[11:cap(expectedFile)]).Build(), nil).Once()
	mockStream.SetupEndOfResponse()
	mockStream.SetupStreamClose()

	response := httptest.NewRecorder()
	context := ginfixture.CreateContext(response);
	request, _ := http.NewRequest(http.MethodPost, expectedEndpoint, strings.NewReader(createValidRequest()))
	context.Request = request

	GeneratePreview(context)

	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, expectedContentType, response.Header().Get(httpheader.ContentType))
	assert.Equal(t, expectedContentLength, response.Header().Get(httpheader.ContentLength))
	assert.Equal(t, expectedContentDisposition, response.Header().Get(httpheader.ContentDisposition))
	assert.EqualValues(t, expectedFile, response.Body.Bytes())
}

func TestGeneratePreview_ErrorOnClose_FileIsStillSent(t *testing.T) {
	closingError := errors.New("error while closing")

	mediaType := "application/pdf"
	size := uint64(134034)
	extension := ".pdf"
	expectedContentType := mediaType
	expectedContentLength := "134034";
	expectedContentDisposition := `attachment; filename="invoice.pdf"`;
	expectedFile := []byte{84,104,101,32,97,110,115,119,101,114,32,105,115,32,52,50}

	mockDocumentServiceGateway := mocks.NewDocumentServiceGateway(t)
	documentServiceGateway = mockDocumentServiceGateway
	mockClient := mocks.NewDocumentServiceClient(t)
	mockStream := mocks.NewDocumentService_GeneratePreviewClient(t)
	mockDocumentServiceGateway.SetupDocumentServiceGatewayToReturnClient(mockClient)
	mockClient.SetupGeneratePreview(mockStream)

	mockStream.EXPECT().Recv().
		Return(fixture.NewGeneratePreviewResponseMetadataBuilder().
			WithMediaType(mediaType).
			WithSize(size).
			WithExtension(extension).
			Build(), nil).Once()
	mockStream.EXPECT().Recv().Return(fixture.NewGeneratePreviewResponseChunkBuilder().WithChunk(expectedFile[0:6]).Build(), nil).Once()
	mockStream.EXPECT().Recv().Return(fixture.NewGeneratePreviewResponseChunkBuilder().WithChunk(expectedFile[6:11]).Build(), nil).Once()
	mockStream.EXPECT().Recv().Return(fixture.NewGeneratePreviewResponseChunkBuilder().WithChunk(expectedFile[11:cap(expectedFile)]).Build(), nil).Once()
	mockStream.SetupEndOfResponse()
	mockStream.EXPECT().CloseSend().Return(closingError).Once()

	response := httptest.NewRecorder()
	context := ginfixture.CreateContext(response);
	request, _ := http.NewRequest(http.MethodPost, expectedEndpoint, strings.NewReader(createValidRequest()))
	context.Request = request

	GeneratePreview(context)

	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, expectedContentType, response.Header().Get(httpheader.ContentType))
	assert.Equal(t, expectedContentLength, response.Header().Get(httpheader.ContentLength))
	assert.Equal(t, expectedContentDisposition, response.Header().Get(httpheader.ContentDisposition))
	assert.EqualValues(t, expectedFile, response.Body.Bytes())
	// TODO Log error
}

func TestGeneratePreview_ErrorWhileConnecting(t *testing.T) {
	connectingError := errors.New("error while receiving")

	mockDocumentServiceGateway := mocks.NewDocumentServiceGateway(t)
	documentServiceGateway = mockDocumentServiceGateway
	mockClient := mocks.NewDocumentServiceClient(t)
	mockDocumentServiceGateway.SetupDocumentServiceGatewayToReturnClient(mockClient)
	mockClient.SetupGeneratePreviewThrowsError(connectingError)

	response := httptest.NewRecorder()
	context := ginfixture.CreateContext(response);
	request, _ := http.NewRequest(http.MethodPost, expectedEndpoint, strings.NewReader(createValidRequest()))
	context.Request = request

	GeneratePreview(context)

	assert.Equal(t, http.StatusServiceUnavailable, response.Code)
}

func TestGeneratePreview_ErrorFromStreamWhileWaitingForMetadata(t *testing.T) {
	receivingError := errors.New("error while receiving")

	mockDocumentServiceGateway := mocks.NewDocumentServiceGateway(t)
	documentServiceGateway = mockDocumentServiceGateway
	mockClient := mocks.NewDocumentServiceClient(t)
	mockStream := mocks.NewDocumentService_GeneratePreviewClient(t)
	mockDocumentServiceGateway.SetupDocumentServiceGatewayToReturnClient(mockClient)
	mockClient.SetupGeneratePreview(mockStream)
	mockStream.EXPECT().Recv().Return(nil, receivingError).Once()

	response := httptest.NewRecorder()
	context := ginfixture.CreateContext(response);
	request, _ := http.NewRequest(http.MethodPost, expectedEndpoint, strings.NewReader(createValidRequest()))
	context.Request = request

	GeneratePreview(context)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
}

func TestGeneratePreview_ErrorFromStreamWhileWaitingForFile(t *testing.T) {
	receivingError := errors.New("error while receiving")

	mockDocumentServiceGateway := mocks.NewDocumentServiceGateway(t)
	documentServiceGateway = mockDocumentServiceGateway
	mockClient := mocks.NewDocumentServiceClient(t)
	mockStream := mocks.NewDocumentService_GeneratePreviewClient(t)
	mockDocumentServiceGateway.SetupDocumentServiceGatewayToReturnClient(mockClient)
	mockClient.SetupGeneratePreview(mockStream)
	mockStream.SetupStreamValidMetadata()
	mockStream.EXPECT().Recv().Return(nil, receivingError).Once()

	response := httptest.NewRecorder()
	context := ginfixture.CreateContext(response);
	request, _ := http.NewRequest(http.MethodPost, expectedEndpoint, strings.NewReader(createValidRequest()))
	context.Request = request

	GeneratePreview(context)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
}

func createValidRequest() string {
	request := `{}`
	return request
}