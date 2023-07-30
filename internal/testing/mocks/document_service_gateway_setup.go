package mocks

import (
	"errors"
)

func (mock *DocumentServiceGateway) SetupDocumentServiceGatewayToReturnDialError() {
	mock.On("CreateDocumentServiceClient").Return(nil, errors.New("i can not reproduce a dial error"))
}

func (mock *DocumentServiceGateway) SetupDocumentServiceGatewayToReturnClient(mockClient *DocumentServiceClient) {
	mock.On("CreateDocumentServiceClient").Return(mockClient, nil)
}