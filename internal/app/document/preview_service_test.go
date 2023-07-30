package document

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kinneko-de/restaurant-document-design-gateway/internal/mocks"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/testfixture"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGeneratePreview_RequestIsNil(t *testing.T) {
	response := httptest.NewRecorder()
	context := testfixture.CreateGinContext(response);
	GeneratePreview(context)
	
	assert.EqualValues(t, http.StatusBadRequest, response.Code)
}

func TestValid(t *testing.T) {
	documentService := &mocks.DocumentServiceServer{}
	documentService.On("GeneratePreview", mock.AnythingOfType("v1.GeneratePreviewRequest"), mock.AnythingOfType("v1.DocumentService_GeneratePreviewServer")).Return(nil)
}