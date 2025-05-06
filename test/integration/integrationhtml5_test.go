package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"web-crawler/internal/api"
	"web-crawler/internal/model"
)

func TestCrawlHTML5Integration(t *testing.T) {
	mockHTML := `
	<!DOCTYPE html>
	<html>
	<head><title>HTML5 Site</title></head>
	<body>
		<header>Header Content</header>
		<img src="logo.png" alt="Logo">
		<footer>Footer Content</footer>
	</body>
	</html>
	`
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, mockHTML)
	}))
	defer mockServer.Close()

	payload := map[string][]string{"urls": {mockServer.URL}}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/start", bytes.NewReader(body))
	w := httptest.NewRecorder()

	api.StartCrawlHandler(w, req)
	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", resp.StatusCode)
	}

	var blocks []model.Block
	if err := json.NewDecoder(resp.Body).Decode(&blocks); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	foundTypes := make(map[string]bool)
	for _, b := range blocks {
		foundTypes[b.Type] = true
	}

	for _, blockType := range []string{"html5_img", "html5_header", "html5_footer"} {
		if !foundTypes[blockType] {
			t.Errorf("Expected block of type %s not found", blockType)
		}
	}
}
