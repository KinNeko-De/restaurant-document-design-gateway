package document

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	apiProtobuf "github.com/kinneko-de/api-contract/golang/kinnekode/protobuf"
	apiRestaurantDocument "github.com/kinneko-de/api-contract/golang/kinnekode/restaurant/document/v1"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/operation"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/timeout"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/httpheader"
	"golang.org/x/time/rate"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	documentServiceGateway DocumentServiceGateway = DocumentServiceGateKeeper{}
	rateLimiters           sync.Map
)

type GeneratePreviewRequest struct {
}

func GeneratePreview(ctx *gin.Context) {
	var request GeneratePreviewRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "request can not be parsed"})
		return
	}
	GeneratePreviewDemo(ctx)
}

func GeneratePreviewDemo(ctx *gin.Context) {
	userId := ctx.Keys["userId"]
	if userId == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user id is not set"})
		return
	}
	if requestIsLimited(userId.(string)) {
		ctx.JSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded. try again later"})
		return
	}

	previewRequest := generateTestDocument()
	err := generatePreview(ctx, previewRequest)
	if err != nil {
		operation.Logger.Error().Err(err).Msg("Failed to generate preview")
		return
	}

	ctx.Status(http.StatusCreated)
}

func requestIsLimited(userId string) bool {
	rateLimiter, _ := rateLimiters.LoadOrStore(userId, createRateLimiter())
	if rateLimiter.(*rate.Limiter).Allow() {
		return false
	} else {
		return true
	}
}

func createRateLimiter() *rate.Limiter {
	rateLimiter := rate.NewLimiter(rate.Every(20*time.Minute), 3)
	return rateLimiter
}

func generatePreview(ctx *gin.Context, previewRequest *apiRestaurantDocument.GeneratePreviewRequest) (err error) {
	fileName := strings.ReplaceAll(uuid.New().String(), "-", "")
	client, err := documentServiceGateway.CreateDocumentServiceClient()
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err)
		return
	}

	callContext, cancelFunc := context.WithDeadline(context.Background(), timeout.GetDeadline(timeout.LongDeadline))
	defer cancelFunc()
	stream, err := client.GeneratePreview(callContext, previewRequest)
	if err != nil {
		ctx.AbortWithError(http.StatusServiceUnavailable, err)
		return
	}

	firstResponse, err := stream.Recv()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	_, ok := firstResponse.File.(*apiRestaurantDocument.GeneratePreviewResponse_Metadata)
	if !ok {
		err = errors.New("FileCase of type 'apidocument.DownloadDocumentResponse_Metadata' expected. Actual value is " + reflect.TypeOf(firstResponse.File).String() + ".")
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var metadata = firstResponse.GetMetadata()
	ctx.Header(httpheader.ContentType, metadata.MediaType)
	ctx.Header(httpheader.ContentLength, strconv.FormatUint(metadata.Size, 10))
	ctx.Header(httpheader.ContentDisposition, fmt.Sprintf("inline; filename=\"%s%s\"", fileName, metadata.Extension))
	ctx.Status(http.StatusCreated)

	for {
		current, done, requestErr := readNextResponse(stream)
		if done {
			break
		}
		if requestErr != nil {
			err = requestErr
			ctx.AbortWithError(http.StatusInternalServerError, requestErr)
			break
		}
		if somethingElseThanChunkWasSent(current) {
			err = fmt.Errorf("FileCase of type 'apiDocumentService.GeneratePreviewResponse_Chunk' expected. Actual value is %s", reflect.TypeOf(current.File).String())
			ctx.AbortWithError(http.StatusInternalServerError, err)
			break
		}

		var chunk = current.GetChunk()
		_, bodyWriteErr := ctx.Writer.Write(chunk)
		if bodyWriteErr != nil {
			err = bodyWriteErr
			ctx.AbortWithError(http.StatusInternalServerError, bodyWriteErr)
			break
		}

	}

	return
}

func generateTestDocument() *apiRestaurantDocument.GeneratePreviewRequest {
	previewRequest := &apiRestaurantDocument.GeneratePreviewRequest{
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
		},
	}
	return previewRequest
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
