package wordpress_test

import (
	"testing"
	"web-crawler/internal/parser/wordpress"

	"github.com/stretchr/testify/assert"
)

func TestMatchImage(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected bool
	}{
		{"Contains wp-image", `<img class="wp-image-123" src="image.jpg">`, true},
		{"Contains wp-post-image", `<img class="wp-post-image" src="image.jpg">`, true},
		{"No wp class", `<img class="some-other-class" src="image.jpg">`, false},
		{"No img tag", `<div class="wp-image-123"></div>`, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			match := wordpress.MatchImage(tt.html)
			assert.Equal(t, tt.expected, match)
		})
	}
}

func TestExtractImage(t *testing.T) {
	t.Run("Extract valid wp-image", func(t *testing.T) {
		html := `<html><body><img class="wp-image-123" src="image.jpg"></body></html>`
		pageURL := "https://test.com"
		block, err := wordpress.ExtractImage(html, pageURL)

		assert.NoError(t, err)
		assert.Equal(t, "wordpress_img", block.Type)
		assert.Contains(t, block.HTML, `wp-image-123`)
		assert.Equal(t, pageURL, block.PageURL)
		assert.Equal(t, "true", block.Found)
	})

	t.Run("No image found", func(t *testing.T) {
		html := `<html><body><p>No image here</p></body></html>`
		pageURL := "https://test.com"
		block, err := wordpress.ExtractImage(html, pageURL)

		assert.Error(t, err)
		assert.Equal(t, "false", block.Found)
		assert.Equal(t, "0.0", block.Accuracy)
	})
}
