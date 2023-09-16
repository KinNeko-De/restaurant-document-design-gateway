package document

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"sync"
	"testing"

	"github.com/kinneko-de/restaurant-document-design-gateway/internal/httpheader"
	fixture "github.com/kinneko-de/restaurant-document-design-gateway/internal/testing/document"
	mocks "github.com/kinneko-de/restaurant-document-design-gateway/internal/testing/document/mocks"
	ginfixture "github.com/kinneko-de/restaurant-document-design-gateway/internal/testing/gin"
	ginmocks "github.com/kinneko-de/restaurant-document-design-gateway/internal/testing/gin/mocks"
	"github.com/stretchr/testify/assert"
	testifymock "github.com/stretchr/testify/mock"
)

const expectedEndpoint string = "/document/preview"

func TestGeneratePreview_RequestIsNil(t *testing.T) {
	response := httptest.NewRecorder()
	context := ginfixture.CreateContext(response)
	GeneratePreview(context)

	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	t.Cleanup(Cleanup)
}

func TestGeneratePreview_DialError(t *testing.T) {
	t.Setenv(hostEnv, "http://localhost")
	t.Setenv(portEnv, "8080")

	mockDocumentServiceGateway := mocks.NewDocumentServiceGateway(t)
	mockDocumentServiceGateway.SetupDocumentServiceGatewayToReturnDialError()
	documentServiceGateway = mockDocumentServiceGateway

	response := httptest.NewRecorder()
	context := ginfixture.CreateContext(response)

	request, _ := http.NewRequest(http.MethodPost, expectedEndpoint, strings.NewReader(fixture.CreateValidGeneratePreviewRequest()))
	context.Request = request

	GeneratePreview(context)

	assert.EqualValues(t, http.StatusServiceUnavailable, response.Code)

	t.Cleanup(Cleanup)
}

func TestGeneratePreview_ValidRequest(t *testing.T) {
	t.Setenv(hostEnv, "http://localhost")
	t.Setenv(portEnv, "8080")

	mediaType := "application/pdf"
	size := uint64(134034)
	extension := ".pdf"
	expectedContentType := mediaType
	expectedContentLength := "134034"
	expectedContentDisposition := `inline; filename="[A-z0-9]+\.pdf"`
	expectedFile := []byte{84, 104, 101, 32, 97, 110, 115, 119, 101, 114, 32, 105, 115, 32, 52, 50}

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
	context := ginfixture.CreateContext(response)
	request, _ := http.NewRequest(http.MethodPost, expectedEndpoint, strings.NewReader(fixture.CreateValidGeneratePreviewRequest()))
	context.Request = request

	GeneratePreview(context)

	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, expectedContentType, response.Header().Get(httpheader.ContentType))
	assert.Equal(t, expectedContentLength, response.Header().Get(httpheader.ContentLength))
	assert.Regexp(t, regexp.MustCompile(expectedContentDisposition), response.Header().Get(httpheader.ContentDisposition))
	assert.EqualValues(t, expectedFile, response.Body.Bytes())

	t.Cleanup(Cleanup)
}

func TestGeneratePreview_ErrorOnClose_FileIsStillSent(t *testing.T) {
	t.Setenv(hostEnv, "http://localhost")
	t.Setenv(portEnv, "8080")

	closingError := errors.New("error while closing")

	mediaType := "application/pdf"
	size := uint64(134034)
	extension := ".pdf"
	expectedFile := []byte{84, 104, 101, 32, 97, 110, 115, 119, 101, 114, 32, 105, 115, 32, 52, 50}

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
	context := ginfixture.CreateContext(response)
	request, _ := http.NewRequest(http.MethodPost, expectedEndpoint, strings.NewReader(fixture.CreateValidGeneratePreviewRequest()))
	context.Request = request

	GeneratePreview(context)

	assert.Equal(t, http.StatusCreated, response.Code)
	assert.EqualValues(t, expectedFile, response.Body.Bytes())
	// TODO Log error

	t.Cleanup(Cleanup)
}

func TestGeneratePreview_ErrorWhileConnecting(t *testing.T) {
	t.Setenv(hostEnv, "http://localhost")
	t.Setenv(portEnv, "8080")

	connectingError := errors.New("error while receiving")

	mockDocumentServiceGateway := mocks.NewDocumentServiceGateway(t)
	documentServiceGateway = mockDocumentServiceGateway
	mockClient := mocks.NewDocumentServiceClient(t)
	mockDocumentServiceGateway.SetupDocumentServiceGatewayToReturnClient(mockClient)
	mockClient.SetupGeneratePreviewThrowsError(connectingError)

	response := httptest.NewRecorder()
	context := ginfixture.CreateContext(response)
	request, _ := http.NewRequest(http.MethodPost, expectedEndpoint, strings.NewReader(fixture.CreateValidGeneratePreviewRequest()))
	context.Request = request

	GeneratePreview(context)

	assert.Equal(t, http.StatusServiceUnavailable, response.Code)

	t.Cleanup(Cleanup)
}

func TestGeneratePreview_ErrorFromStreamWhileWaitingForMetadata(t *testing.T) {
	t.Setenv(hostEnv, "http://localhost")
	t.Setenv(portEnv, "8080")

	receivingError := errors.New("error while receiving")

	mockDocumentServiceGateway := mocks.NewDocumentServiceGateway(t)
	documentServiceGateway = mockDocumentServiceGateway
	mockClient := mocks.NewDocumentServiceClient(t)
	mockStream := mocks.NewDocumentService_GeneratePreviewClient(t)
	mockDocumentServiceGateway.SetupDocumentServiceGatewayToReturnClient(mockClient)
	mockClient.SetupGeneratePreview(mockStream)
	mockStream.EXPECT().Recv().Return(nil, receivingError).Once()

	response := httptest.NewRecorder()
	context := ginfixture.CreateContext(response)
	request, _ := http.NewRequest(http.MethodPost, expectedEndpoint, strings.NewReader(fixture.CreateValidGeneratePreviewRequest()))
	context.Request = request

	GeneratePreview(context)

	assert.Equal(t, http.StatusInternalServerError, response.Code)

	t.Cleanup(Cleanup)
}

func TestGeneratePreview_ErrorFromStreamWhileWaitingForFile(t *testing.T) {
	t.Setenv(hostEnv, "http://localhost")
	t.Setenv(portEnv, "8080")

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
	context := ginfixture.CreateContext(response)
	request, _ := http.NewRequest(http.MethodPost, expectedEndpoint, strings.NewReader(fixture.CreateValidGeneratePreviewRequest()))
	context.Request = request

	GeneratePreview(context)

	assert.Equal(t, http.StatusInternalServerError, response.Code)

	t.Cleanup(Cleanup)
}

func TestGeneratePreview_ChunkSentBeforeMetadata(t *testing.T) {
	t.Setenv(hostEnv, "http://localhost")
	t.Setenv(portEnv, "8080")

	mockDocumentServiceGateway := mocks.NewDocumentServiceGateway(t)
	documentServiceGateway = mockDocumentServiceGateway
	mockClient := mocks.NewDocumentServiceClient(t)
	mockStream := mocks.NewDocumentService_GeneratePreviewClient(t)
	mockDocumentServiceGateway.SetupDocumentServiceGatewayToReturnClient(mockClient)
	mockClient.SetupGeneratePreview(mockStream)
	mockStream.SetupStreamValidChunk()

	response := httptest.NewRecorder()
	context := ginfixture.CreateContext(response)
	request, _ := http.NewRequest(http.MethodPost, expectedEndpoint, strings.NewReader(fixture.CreateValidGeneratePreviewRequest()))
	context.Request = request

	GeneratePreview(context)

	assert.Equal(t, http.StatusInternalServerError, response.Code)

	t.Cleanup(Cleanup)
}

func TestGeneratePreview_MetadataIsSentTwice(t *testing.T) {
	t.Setenv(hostEnv, "http://localhost")
	t.Setenv(portEnv, "8080")

	mockDocumentServiceGateway := mocks.NewDocumentServiceGateway(t)
	documentServiceGateway = mockDocumentServiceGateway
	mockClient := mocks.NewDocumentServiceClient(t)
	mockStream := mocks.NewDocumentService_GeneratePreviewClient(t)
	mockDocumentServiceGateway.SetupDocumentServiceGatewayToReturnClient(mockClient)
	mockClient.SetupGeneratePreview(mockStream)
	mockStream.SetupStreamValidMetadata()
	mockStream.SetupStreamValidMetadata()

	response := httptest.NewRecorder()
	context := ginfixture.CreateContext(response)
	request, _ := http.NewRequest(http.MethodPost, expectedEndpoint, strings.NewReader(fixture.CreateValidGeneratePreviewRequest()))
	context.Request = request

	GeneratePreview(context)

	assert.Equal(t, http.StatusInternalServerError, response.Code)

	t.Cleanup(Cleanup)
}

func TestGeneratePreview_HttpContextWriterError(t *testing.T) {
	t.Setenv(hostEnv, "http://localhost")
	t.Setenv(portEnv, "8080")

	writerError := errors.New("error while writing into http respnse")

	mockDocumentServiceGateway := mocks.NewDocumentServiceGateway(t)
	documentServiceGateway = mockDocumentServiceGateway
	mockClient := mocks.NewDocumentServiceClient(t)
	mockStream := mocks.NewDocumentService_GeneratePreviewClient(t)
	mockDocumentServiceGateway.SetupDocumentServiceGatewayToReturnClient(mockClient)
	mockClient.SetupGeneratePreview(mockStream)
	mockStream.SetupStreamValidMetadata()
	mockStream.SetupStreamValidChunk()

	response := httptest.NewRecorder()
	context := ginfixture.CreateContext(response)
	writerMock := ginmocks.CreateResponseWriterMock(t, response)
	writerMock.SetupWriteError(writerError)
	context.Writer = writerMock

	request, _ := http.NewRequest(http.MethodPost, expectedEndpoint, strings.NewReader(fixture.CreateValidGeneratePreviewRequest()))
	context.Request = request

	GeneratePreview(context)

	assert.Equal(t, http.StatusInternalServerError, response.Code)

	t.Cleanup(Cleanup)
}

func TestGeneratePreview_Unauthorized(t *testing.T) {
	t.Setenv(hostEnv, "http://localhost")
	t.Setenv(portEnv, "8080")

	response := httptest.NewRecorder()
	context := ginfixture.CreateContext(response)
	// TODO replace with clear function in Go 1.21
	for k := range context.Keys {
		delete(context.Keys, k)
	}
	request, _ := http.NewRequest(http.MethodPost, expectedEndpoint, strings.NewReader(fixture.CreateValidGeneratePreviewRequest()))
	context.Request = request

	GeneratePreview(context)

	assert.Equal(t, http.StatusUnauthorized, response.Code)
}

func TestGeneratePreview_IsRateLimited(t *testing.T) {
	limitedAfter := int(4)
	t.Setenv(hostEnv, "http://localhost")
	t.Setenv(portEnv, "8080")

	var lastResponse *httptest.ResponseRecorder
	for i := 0; i < limitedAfter; i++ {
		mediaType := "application/pdf"
		size := uint64(134034)
		extension := ".pdf"
		expectedFile := []byte{84, 104, 101, 32, 97, 110, 115, 119, 101, 114, 32, 105, 115, 32, 52, 50}

		mockDocumentServiceGateway := mocks.NewDocumentServiceGateway(t)
		documentServiceGateway = mockDocumentServiceGateway
		mockClient := mocks.NewDocumentServiceClient(t)
		mockStream := mocks.NewDocumentService_GeneratePreviewClient(t)
		mockDocumentServiceGateway.EXPECT().CreateDocumentServiceClient().Return(mockClient, nil).Maybe()
		mockClient.EXPECT().GeneratePreview(testifymock.Anything, testifymock.Anything).Return(mockStream, nil).Maybe()
		mockStream.EXPECT().Recv().
			Return(fixture.NewGeneratePreviewResponseMetadataBuilder().
				WithMediaType(mediaType).
				WithSize(size).
				WithExtension(extension).
				Build(), nil).Maybe()
		mockStream.EXPECT().Recv().Return(fixture.NewGeneratePreviewResponseChunkBuilder().WithChunk(expectedFile[0:6]).Build(), nil).Maybe()
		mockStream.EXPECT().Recv().Return(fixture.NewGeneratePreviewResponseChunkBuilder().WithChunk(expectedFile[6:11]).Build(), nil).Maybe()
		mockStream.EXPECT().Recv().Return(fixture.NewGeneratePreviewResponseChunkBuilder().WithChunk(expectedFile[11:cap(expectedFile)]).Build(), nil).Maybe()

		mockStream.EXPECT().Recv().Return(nil, io.EOF).Maybe()
		mockStream.EXPECT().CloseSend().Return(nil).Maybe()

		response := httptest.NewRecorder()
		context := ginfixture.CreateContext(response)
		request, _ := http.NewRequest(http.MethodPost, expectedEndpoint, strings.NewReader(fixture.CreateValidGeneratePreviewRequest()))
		context.Request = request

		GeneratePreview(context)

		lastResponse = response
	}

	assert.Equal(t, http.StatusTooManyRequests, lastResponse.Code)

	t.Cleanup(Cleanup)
}

func Cleanup() {
	rateLimiters = sync.Map{}
}
