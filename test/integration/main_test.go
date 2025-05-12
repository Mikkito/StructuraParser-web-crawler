package integration_test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"web-crawler/pkg/utils/logger"
	"web-crawler/test/integration/testutils"
)

var mockServer *httptest.Server

func TestMain(m *testing.M) {
	_ = os.Chdir("../..")
	err := logger.Init("pkg/utils/logger/config.yaml")
	if err != nil {
		fmt.Printf("logger init failed: %v\n", err)
		os.Exit(1)
	}
	mux := http.NewServeMux()

	mockDir := "test/integration/mockdata"

	files, err := os.ReadDir(mockDir)
	if err != nil {
		fmt.Printf("failed to read mockdata directory: %v\n", err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".html") {
			continue
		}

		name := file.Name()
		route := "/" + strings.TrimSuffix(name, ".html")

		html := testutils.LoadMockHTML(nil, name)

		mux.HandleFunc(route, func(html string) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/html")
				w.Write([]byte(html))
			}
		}(html))
	}

	mockServer = httptest.NewServer(mux)
	defer mockServer.Close()

	code := m.Run()
	mockServer.Close()
	os.Exit(code)
}

type nilSafeT struct{}

func (n nilSafeT) Helper()                           {}
func (n nilSafeT) Fatalf(format string, args ...any) { log.Fatalf(format, args...) }
