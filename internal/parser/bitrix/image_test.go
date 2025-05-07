package bitrix_test

import (
	"testing"
	"web-crawler/internal/parser/bitrix"

	"github.com/stretchr/testify/assert"
)

func TestMatchImage(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected bool
	}{
		{"Contains bx-image", `<img class="bx-image" src="image.jpg">`, true},
		{"Contains bitrix-image", `<img class="bitrix-image" src="image.jpg">`, true},
		{"Other class", `<img class="not-bitrix" src="image.jpg">`, false},
		{"No img tag", `<div class="bx-image"></div>`, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			match := bitrix.MatchImage(tt.html)
			assert.Equal(t, tt.expected, match)
		})
	}
}

func TestExtractImage(t *testing.T) {
	t.Run("Extract valid bitrix image", func(t *testing.T) {
		html := `<html><body><img class="bx-image" src="bitrix.jpg"></body></html>`
		pageURL := "https://bitrix.example"
		block, err := bitrix.ExtractImage(html, pageURL)

		assert.NoError(t, err)
		assert.Equal(t, "bitrix_image", block.Type)
		assert.Contains(t, block.HTML, "bx-image")
		assert.Equal(t, "true", block.Found)
		assert.Equal(t, pageURL, block.PageURL)
	})

	t.Run("No image found", func(t *testing.T) {
		html := `<html><body><p>No image</p></body></html>`
		pageURL := "https://bitrix.test"
		block, err := bitrix.ExtractImage(html, pageURL)

		assert.Error(t, err)
		assert.Equal(t, "false", block.Found)
		assert.Equal(t, "0.0", block.Accuracy)
	})
}
