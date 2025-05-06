package testutils

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// LoadMockHTML загружает HTML-мок из файла
func LoadMockHTML(t *testing.T, filename string) string {
	if t != nil {
		t.Helper()
	}
	path := filepath.Join("tests", "integration", "mockdata", filename)
	data, err := os.ReadFile(path)
	if err != nil {
		if t != nil {
			t.Fatalf("failed to load mock HTML %s: %v", filename, err)
		} else {
			panic(fmt.Sprintf("failed to load mock HTML %s: %v", filename, err))
		}
	}
	return string(data)
}
