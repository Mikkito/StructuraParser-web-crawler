package crawler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"web-crawler/internal/crawler"
	"web-crawler/internal/dispatcher"
	"web-crawler/internal/model"

	"go.uber.org/zap"
)

type mockTildaHeaderHandler struct{}

func (m *mockTildaHeaderHandler) Type() string { return "tilda_header" }

func (m *mockTildaHeaderHandler) Match(html string) bool {
	return true
}

func (m *mockTildaHeaderHandler) Extract(html, url string) (model.Block, error) {
	return model.Block{
		Type:     "tilda_header",
		HTML:     "<header class='t-header'>Mock Header</header>",
		PageURL:  url,
		Found:    "true",
		Accuracy: "0.9",
	}, nil
}

func TestScrapeURL_Success(t *testing.T) {
	// registr handler mock
	dispatcher.ResetHandlers()
	model.RegisterHandler(&mockTildaHeaderHandler{})

	// Server mock
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<html><header class="t-header">Test</header></html>`))
	}))
	defer server.Close()

	resultChan := make(chan model.Block, 1)

	log := zap.NewNop().Sugar()

	// Run scrapeURL
	crawler.ScrapeURL(server.URL, resultChan, log)

	// Result check
	select {
	case block := <-resultChan:
		if block.Type != "tilda_header" {
			t.Errorf("Expected type tilda_header, got %s", block.Type)
		}
		if block.PageURL != server.URL {
			t.Errorf("Expected URL %s, got %s", server.URL, block.PageURL)
		}
	default:
		t.Fatal("No block was returned")
	}
}
