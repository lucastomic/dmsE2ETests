package e2e

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func helperCreateMultipartRequest(
	t *testing.T,
	method, url, fileName, fileContent string,
	headers map[string]string,
	otherFields map[string]string,
) (*http.Request, *bytes.Buffer) {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	if fileName != "" && fileContent != "" {
		part, err := writer.CreateFormFile("uploadFile", fileName)
		if err != nil {
			t.Fatalf("Failed to create form file: %v", err)
		}
		_, err = part.Write([]byte(fileContent))
		if err != nil {
			t.Fatalf("Failed to write to form file: %v", err)
		}
	}

	for k, v := range otherFields {
		writer.WriteField(k, v)
	}

	err := writer.Close()
	if err != nil {
		t.Fatalf("Failed to close writer: %v", err)
	}

	req, err := http.NewRequest(method, url, &buffer)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return req, &buffer
}

func helperCreateGETRequest(
	t *testing.T,
	url string,
	headers map[string]string,
	queryParams []string,
) (*http.Request, *bytes.Buffer) {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	err := writer.Close()
	if err != nil {
		t.Fatalf("Failed to close writer: %v", err)
	}
	for _, q := range queryParams {
		url = url + "/" + q
	}

	req, err := http.NewRequest("GET", url, &buffer)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return req, &buffer
}

func helperSendRequest(t *testing.T, req *http.Request, expectedStatusCode int) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, expectedStatusCode, resp.StatusCode, "Unexpected status code")
}
