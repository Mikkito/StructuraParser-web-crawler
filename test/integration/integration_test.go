package integration_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"web-crawler/internal/api"
	_ "web-crawler/internal/handlers"
	"web-crawler/internal/model"
	"web-crawler/test/integration/testutils"
)

func TestCrawlWithMockSite(t *testing.T) {

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		html := testutils.LoadMockHTML(t, "automatisation.html")
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
	req := httptest.NewRequest(http.MethodPost, "/crawl", bytes.NewBuffer(body))
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
	knownTypes := map[string]bool{
		"html5_header":     true,
		"html5_footer":     true,
		"html5_img":        true,
		"bitrix_header":    true,
		"bitrix_footer":    true,
		"bitrix_img":       true,
		"tilda_block":      true,
		"wordpress_header": true,
		"wordpress_footer": true,
	}
	foundTypes := make(map[string]int)
	for _, b := range blocks {
		if b.Type == "" || b.HTML == "" || b.PageURL == "" {
			t.Errorf("Incomplete block: %+v", b)
		}
		if !knownTypes[b.Type] {
			t.Errorf("Unknown block type found: %s", b.Type)
		}
		foundTypes[b.Type]++
	}
	t.Logf("Found block types: %+v", foundTypes)

	// Можно проверить наличие конкретных типов
	//if foundTypes["html5_header"] == 0 {
	//	t.Errorf("Expected at least one html5_header block")
	//}
}
