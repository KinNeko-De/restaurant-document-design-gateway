package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePreview(t *testing.T) {
	t.Skip("test is not working yet.")
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/document/preview", nil)
	router.ServeHTTP(w, req)

	_ = assert.Equal(t, http.StatusCreated, w.Code)


	/*
	var response = w.Result()

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Errorf("Result can not be closed: %e.", err)
		}
	}(response.Body)

	type CreateOrderResponse struct {
		Id uuid.UUID
	}

	actualResponse := readResponse[CreateOrderResponse](t, response)

	assert.NotNil(t, actualResponse.Id)
	assert.IsType(t, uuid.UUID{}, actualResponse.Id)
	*/
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
