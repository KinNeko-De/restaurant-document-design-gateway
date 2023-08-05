package mocks

import (
	"errors"
)

func (mock *DocumentServiceGateway) SetupDocumentServiceGatewayToReturnDialError() {
	mock.EXPECT().CreateDocumentServiceClient().Return(nil, errors.New("i can not reproduce a dial error"))
}

func (mock *DocumentServiceGateway) SetupDocumentServiceGatewayToReturnClient(mockClient *DocumentServiceClient) {
	mock.EXPECT().CreateDocumentServiceClient().Return(mockClient, nil)
}