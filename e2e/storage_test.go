package e2e

import (
	"math/rand"
	"net/http"
	"strconv"
	"testing"
)

var storageURL = "http://localhost:3001/storage/file"

func TestStorageUploadSuccessful(t *testing.T) {
	req, _ := helperCreateMultipartRequest(
		t,
		"POST",
		storageURL,
		"test.txt",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		map[string]string{},
		// TODO: Instad of using a random number, we should have predefined a constant number, which is realised after testing for future testing
		map[string]string{"Id": strconv.Itoa(rand.Intn(1000) + 1000)},
	)

	helperSendRequest(t, req, 201)
}

func TestStorageUploadWithoutID(t *testing.T) {
	req, _ := helperCreateMultipartRequest(
		t,
		"POST",
		storageURL,
		"test.txt",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		map[string]string{},
		// TODO: Instad of using a random number, we should have predefined a constant number, which is realised after testing for future testing
		map[string]string{},
	)

	helperSendRequest(t, req, 400)
}

func TestRetrieveFile(t *testing.T) {
	setUpTestFileInID1(
		t,
	) // This configuration should not be left like this, but the file should always be configured on the server. It is left because it is believed that this is outside the scope of the task.
	req, _ := helperCreateGETRequest(
		t,
		storageURL,
		map[string]string{},
		[]string{"1"}, // Predefined ID for testing
	)

	helperSendRequest(t, req, 200)
}

func TestRetrieveUnexistentFile(t *testing.T) {
	req, _ := helperCreateGETRequest(
		t,
		storageURL,
		map[string]string{},
		[]string{"-1"},
	)
	helperSendRequest(t, req, 404)
}

func setUpTestFileInID1(t *testing.T) {
	req, _ := helperCreateMultipartRequest(
		t,
		"POST",
		storageURL,
		"test.txt",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		map[string]string{},
		map[string]string{"Id": "1"},
	)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()
}
