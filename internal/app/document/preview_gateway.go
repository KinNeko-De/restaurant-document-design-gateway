package document

import (
	grpccontext "context"
	"errors"
	"io"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	apiProtobuf "github.com/kinneko-de/api-contract/golang/kinnekode/protobuf"
	apiDocumentService "github.com/kinneko-de/api-contract/golang/kinnekode/restaurant/document/v1"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/timeout"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	ApiDocumentService string = "SetByMain"
)

type GeneratePreviewRequest struct {
	RequestId string `json:"requestId"`
}

func GeneratePreview(context *gin.Context) {
	var request GeneratePreviewRequest
	if err := context.BindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "request can not be parsed"})
		return
	}
	requestId, err := uuid.Parse(request.RequestId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "requestId '" + request.RequestId + "' is not a valid uuid. expect uuid in the following format: 550e8400-e29b-11d4-a716-446655440000"})
		return
	}

	c, dialError := grpc.Dial(ApiDocumentService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if dialError != nil {
		context.JSON(http.StatusServiceUnavailable, gin.H{"error": "service not available. please try again later"})
		return
	}

	client := apiDocumentService.NewDocumentServiceClient(c)
	requestUuid, _ := apiProtobuf.ToProtobuf(requestId)
	previewRequest := apiDocumentService.GeneratePreviewRequest{
		RequestId: requestUuid,

	}

	callContext, cancelFunc := grpccontext.WithDeadline(grpccontext.Background(), timeout.GetDefaultDeadline())
	defer cancelFunc()

	stream, clientErr := client.GeneratePreview(callContext, &previewRequest)
	if clientErr != nil {
		context.AbortWithError(http.StatusServiceUnavailable, clientErr)
		return
	}

	firstResponse, streamErr := stream.Recv()
	if streamErr != nil {
		context.AbortWithError(http.StatusServiceUnavailable, streamErr)
		return
	}

	_, ok := firstResponse.File.(*apiDocumentService.GeneratePreviewResponse_Metadata)
	if !ok {
		context.AbortWithError(http.StatusInternalServerError, errors.New("FileCase of type 'apidocument.DownloadDocumentResponse_Metadata' expected. Actual value is "+reflect.TypeOf(firstResponse.File).String()+"."))
		return
	}
	var metadata = firstResponse.GetMetadata()
	context.Header("Content-Type", metadata.MediaType)
	context.Header("Content-Length", strconv.FormatUint(metadata.Size, 10))
	context.Status(http.StatusCreated)

	for {
		current, done, requestErr := readNextResponse(stream)
		if done {
			return
		}
		if requestErr != nil {
			context.AbortWithError(http.StatusServiceUnavailable, requestErr)
			return
		}
		if somethingElseThanChunkWasSent(current) {
			context.AbortWithError(http.StatusInternalServerError, errors.New("FileCase of type 'apiDocumentService.GeneratePreviewResponse_Chunk' expected. Actual value is "+reflect.TypeOf(current.File).String()+"."))
			return
		}

		var chunk = current.GetChunk()
		_, bodyWriteErr := context.Writer.Write(chunk)
		if bodyWriteErr != nil {
			context.AbortWithError(http.StatusInternalServerError, bodyWriteErr)
			return
		}
	}
}

func somethingElseThanChunkWasSent(current *apiDocumentService.GeneratePreviewResponse) bool {
	_, ok := current.File.(*apiDocumentService.GeneratePreviewResponse_Chunk)
	return !ok
}

func readNextResponse(stream apiDocumentService.DocumentService_GeneratePreviewClient) (*apiDocumentService.GeneratePreviewResponse, bool, error) {
	current, err := stream.Recv()

	if err == io.EOF {
		err := stream.CloseSend()
		if err != nil {
			return nil, true, err
		}
		return nil, true, nil
	}

	if err != nil {
		return nil, false, err
	}
	return current, false, nil
}
