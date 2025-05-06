package integration_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"web-crawler/internal/api"
	"web-crawler/internal/model"
	"web-crawler/test/integration/testutils"
)

func TestCrawlWithMockSite(t *testing.T) {

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		html := testutils.LoadMockHTML(t, "tilda.html")
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(html))
	}))
	defer mockServer.Close()

	// POST request
	payload := map[string][]string{
		"urls": {mockServer.URL},
	}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/start", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Run handler
	api.StartCrawlHandler(w, req)
	resp := w.Result()
	defer resp.Body.Close()

	// Check Status
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", resp.StatusCode)
	}

	// Parse JSON-answer
	var blocks []model.Block
	if err := json.NewDecoder(resp.Body).Decode(&blocks); err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}

	// Block check
	if len(blocks) == 0 {
		t.Errorf("Expected at least one block, got 0")
	}
	for _, b := range blocks {
		if b.Type == "" || b.HTML == "" || b.PageURL == "" {
			t.Errorf("Incomplete block: %+v", b)
		}
	}
}
