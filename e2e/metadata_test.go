package e2e

import (
	"testing"
)

func TestMetadataSuccessful(t *testing.T) {
	req, _ := helperCreateMultipartRequest(
		t,
		"GET",
		"http://localhost:3001/metadata/file",
		"test.txt",
		"contenido del archivo",
		map[string]string{},
		map[string]string{},
	)
	helperSendRequest(t, req, 200)
}

func TestMetadataWithWrongMethod(t *testing.T) {
	req, _ := helperCreateMultipartRequest(
		t,
		"POST",
		"http://localhost:3001/metadata/file",
		"test.txt",
		"contenido del archivo",
		map[string]string{},
		map[string]string{},
	)
	helperSendRequest(t, req, 405)
}

func TestMetadataWithoutImage(t *testing.T) {
	req, _ := helperCreateMultipartRequest(
		t,
		"GET",
		"http://localhost:3001/metadata/file",
		"",
		"",
		map[string]string{},
		map[string]string{},
	)
	helperSendRequest(
		t,
		req,
		400,
	)
}
