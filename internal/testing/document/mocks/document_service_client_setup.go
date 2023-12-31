package mocks

import (
	testifymock "github.com/stretchr/testify/mock"
)

func (mock *DocumentServiceClient) SetupGeneratePreview(mockStream *DocumentService_GeneratePreviewClient) {
	mock.EXPECT().GeneratePreview(
		testifymock.Anything,
		testifymock.Anything,
	).Return(mockStream, nil).Once()
}

func (mock *DocumentServiceClient) SetupGeneratePreviewThrowsError(err error) {
	mock.EXPECT().GeneratePreview(
		testifymock.Anything,
		testifymock.Anything,
	).Return(nil, err).Once()
}
