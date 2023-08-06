package main

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	mainfixture "github.com/kinneko-de/restaurant-document-design-gateway/internal/testing/main"
)

const expectedEndpoint string = "/document/preview"

// tests that the service is mapped to the expected expectedEndpoint (no code coverage)
func TestDocumentPreview_RequestIsNil(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, expectedEndpoint, nil)

	response := mainfixture.SendRequestToSut(setupRouter(), request)

	assert.Equal(t, http.StatusBadRequest, response.Code)
}


func readResponse[K any](t *testing.T, response *http.Response) K {
	data := ReadAllBytes(t, response)
	actualResponse := decodeToJson[K](t, data)
	return actualResponse
}

func ReadAllBytes(t *testing.T, response *http.Response) []byte {
	data, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Unable to read response body %v", err)
	}
	return data
}

func decodeToJson[K any](t *testing.T, data []byte) K {
	var actualResponse K
	err := json.Unmarshal(
		data,
		&actualResponse,
	)
	if err != nil {
		str1 := string(data[:])
		t.Errorf("Response can not be read to expected response %v. Raw: %s", err, str1)
	}
	return actualResponse
}

func createRequest(requestIdParameter string, requestIdValue string) string {
	request := `{
  "` + requestIdParameter + `": "` + requestIdValue + `"
}`
	return request
}
