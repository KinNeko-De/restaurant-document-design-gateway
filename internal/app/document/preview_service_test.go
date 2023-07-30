package document

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kinneko-de/restaurant-document-design-gateway/internal/testing/ginfixture"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePreview_RequestIsNil(t *testing.T) {
	response := httptest.NewRecorder()
	context := ginfixture.CreateContext(response);
	GeneratePreview(context)
	
	assert.EqualValues(t, http.StatusBadRequest, response.Code)
}

/*
func TestValid(t *testing.T) {
	documentService := &mocks.DocumentServiceServer{}
	documentService.On("GeneratePreview", mock.AnythingOfType("v1.GeneratePreviewRequest"), mock.AnythingOfType("v1.DocumentService_GeneratePreviewServer")).Return(nil)
	
}
*/