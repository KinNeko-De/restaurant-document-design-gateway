package document

import (
	grpccontext "context"
	"errors"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	apiProtobuf "github.com/kinneko-de/api-contract/golang/kinnekode/protobuf"
	apiRestaurantDocument "github.com/kinneko-de/api-contract/golang/kinnekode/restaurant/document/v1"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/timeout"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	client := apiRestaurantDocument.NewDocumentServiceClient(c)
	requestUuid, _ := apiProtobuf.ToProtobuf(requestId)
	previewRequest := apiRestaurantDocument.GeneratePreviewRequest{
		RequestId: requestUuid,
		RequestedDocument: &apiRestaurantDocument.RequestedDocument{
			Type: &apiRestaurantDocument.RequestedDocument_Invoice{
				Invoice: &apiRestaurantDocument.Invoice{
					DeliveredOn:  timestamppb.New(time.Date(2020, time.April, 13, 0, 0, 0, 0, time.UTC)),
					CurrencyCode: "EUR",
					Recipient: &apiRestaurantDocument.Invoice_Recipient{
						Name:     "Max Mustermann",
						Street:   "Musterstraße 17",
						City:     "Musterstadt",
						PostCode: "12345",
						Country:  "DE",
					},
					Items: []*apiRestaurantDocument.Invoice_Item{
						{
							Description: "Spitzenunterwäsche\r\nANS 23054303053",
							Quantity:    2,
							NetAmount:   &apiProtobuf.Decimal{Value: "3.35"},
							Taxation:    &apiProtobuf.Decimal{Value: "19"},
							TotalAmount: &apiProtobuf.Decimal{Value: "3.99"},
							Sum:         &apiProtobuf.Decimal{Value: "7.98"},
						},
						{
							Description: "Schlabberhose (10% reduziert)\r\nANS 606406540",
							Quantity:    1,
							NetAmount:   &apiProtobuf.Decimal{Value: "9.07"},
							Taxation:    &apiProtobuf.Decimal{Value: "19"},
							TotalAmount: &apiProtobuf.Decimal{Value: "10.79"},
							Sum:         &apiProtobuf.Decimal{Value: "10.79"},
						},
						{
							Description: "Versandkosten",
							Quantity:    1,
							NetAmount:   &apiProtobuf.Decimal{Value: "0.00"},
							Taxation:    &apiProtobuf.Decimal{Value: "0"},
							TotalAmount: &apiProtobuf.Decimal{Value: "0.00"},
							Sum:         &apiProtobuf.Decimal{Value: "0.00"},
						},
					},
				},
			},
			OutputFormats: []apiRestaurantDocument.RequestedDocument_OutputFormat{
				apiRestaurantDocument.RequestedDocument_OUTPUT_FORMAT_PDF,
			},
		},
	}

	callContext, cancelFunc := grpccontext.WithDeadline(grpccontext.Background(), timeout.GetDeadline(timeout.LongDeadline))
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

	_, ok := firstResponse.File.(*apiRestaurantDocument.GeneratePreviewResponse_Metadata)
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

func somethingElseThanChunkWasSent(current *apiRestaurantDocument.GeneratePreviewResponse) bool {
	_, ok := current.File.(*apiRestaurantDocument.GeneratePreviewResponse_Chunk)
	return !ok
}

func readNextResponse(stream apiRestaurantDocument.DocumentService_GeneratePreviewClient) (*apiRestaurantDocument.GeneratePreviewResponse, bool, error) {
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
