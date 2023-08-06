package document

import (
	v1 "github.com/kinneko-de/api-contract/golang/kinnekode/restaurant/document/v1"
)

type generatePreviewResponseChunkBuilder struct {
	response *v1.GeneratePreviewResponse
}

func NewGeneratePreviewResponseChunkBuilder() *generatePreviewResponseChunkBuilder {
	chunk := []byte{84,104,101,32,97,110,115,119,101,114,32,105,115,32,52,50}

	return &generatePreviewResponseChunkBuilder{
		response: &v1.GeneratePreviewResponse {
			File: &v1.GeneratePreviewResponse_Chunk{
				Chunk: chunk,
			},
		},
	}
}

func (builder *generatePreviewResponseChunkBuilder) WithChunk(chunk []byte) *generatePreviewResponseChunkBuilder {
	builder.response.File = &v1.GeneratePreviewResponse_Chunk{
		Chunk: chunk,
	}
	return builder
}

func (builder *generatePreviewResponseChunkBuilder) Build() *v1.GeneratePreviewResponse {
	return builder.response
}