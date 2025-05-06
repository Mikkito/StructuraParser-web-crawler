package integration_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"web-crawler/pkg/utils/logger"
	"web-crawler/test/integration/testutils"
)

var mockServer *httptest.Server

func TestMain(m *testing.M) {
	logger.Init("pkg/utils/logger/config.yaml")

	mux := http.NewServeMux()

	mux.HandleFunc("/bitrix", func(w http.ResponseWriter, r *http.Request) {
		html := testutils.LoadMockHTML(nil, "bitrix.html")
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
	})

	mockServer = httptest.NewServer(mux)
	defer mockServer.Close()

	code := m.Run()
	mockServer.Close()
	os.Exit(code)
}

// nilSafeT реализует testing.TB, но ничего не делает
// Используется, если LoadMockHTML вызывается вне тестов
type nilSafeT struct{}

func (n nilSafeT) Helper()                           {}
func (n nilSafeT) Fatalf(format string, args ...any) { log.Fatalf(format, args...) }
