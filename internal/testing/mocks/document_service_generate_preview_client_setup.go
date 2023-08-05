package mocks

import (
	"io"

	v1 "github.com/kinneko-de/api-contract/golang/kinnekode/restaurant/document/v1"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/testing/document"
)


func (mock *DocumentService_GeneratePreviewClient) SetupStreamValidChunk() {
	response := document.NewGeneratePreviewResponseChunkBuilder().Build()
	mock.SetupStreamResponse(response)
}

func (mock *DocumentService_GeneratePreviewClient) SetupStreamValidMetadata() {
	response := document.NewGeneratePreviewResponseMetadataBuilder().Build()
	mock.SetupStreamResponse(response)
}

func (mock *DocumentService_GeneratePreviewClient) SetupStreamResponse(response *v1.GeneratePreviewResponse) {
	mock.EXPECT().Recv().Return(response, nil).Once()
}

func (mock *DocumentService_GeneratePreviewClient) SetupEndOfResponse() {
	mock.EXPECT().Recv().Return(nil, io.EOF).Once()
}


func (mock *DocumentService_GeneratePreviewClient) SetupStreamClose() {
	mock.EXPECT().CloseSend().Return(nil).Once()
}

