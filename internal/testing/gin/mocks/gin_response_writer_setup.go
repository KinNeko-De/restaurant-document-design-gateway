package mocks

import (
	http "net/http"
	"net/http/httptest"
	"testing"

	testifymock "github.com/stretchr/testify/mock"
)

func CreateResponseWriterMock(t *testing.T, response *httptest.ResponseRecorder) *ResponseWriter {
	writerMock := NewResponseWriter(t)

	writerMock.EXPECT().Header().Return(http.Header{})
	writerMock.EXPECT().WriteHeader(testifymock.Anything).Run(
		func(statusCode int) {
			response.Code = statusCode
		},
	)
	writerMock.EXPECT().WriteHeaderNow()
	
	return writerMock
}

func (mock *ResponseWriter) SetupWriteError(err error) {
	mock.EXPECT().Write(testifymock.Anything).Return(0, err)
}
