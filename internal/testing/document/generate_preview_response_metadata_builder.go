package document

import (
	v1 "github.com/kinneko-de/api-contract/golang/kinnekode/restaurant/document/v1"
)

type generatePreviewResponseMetadataBuilder struct {
	response *v1.GeneratePreviewResponse
}

func NewGeneratePreviewResponseMetadataBuilder() *generatePreviewResponseMetadataBuilder {
	mediaType := "application/pdf"
	size := uint64(134034)
	extension := ".pdf"

	return &generatePreviewResponseMetadataBuilder{
		response: &v1.GeneratePreviewResponse {
			File: &v1.GeneratePreviewResponse_Metadata{ 
				Metadata: &v1.GeneratedFileMetadata{ 
					MediaType: mediaType, 
					Size: size, 
					Extension: extension,
				},
			},
		},
	}
}

func (builder *generatePreviewResponseMetadataBuilder) WithMediaType(mediaType string) *generatePreviewResponseMetadataBuilder {
	builder.response.GetMetadata().MediaType = mediaType
	return builder
}

func (builder *generatePreviewResponseMetadataBuilder) WithSize(size uint64) *generatePreviewResponseMetadataBuilder {
	builder.response.GetMetadata().Size = size
	return builder
}

func (builder *generatePreviewResponseMetadataBuilder) WithExtension(extension string) *generatePreviewResponseMetadataBuilder {
	builder.response.GetMetadata().Extension = extension
	return builder
}

func (builder *generatePreviewResponseMetadataBuilder) Build() *v1.GeneratePreviewResponse {
	return builder.response
}