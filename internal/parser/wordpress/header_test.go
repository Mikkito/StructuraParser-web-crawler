package wordpress_test

import (
	"testing"
	"web-crawler/internal/parser/wordpress"
)

func TestMatchHeader(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected bool
	}{
		// Positive test
		{"Contains wp-content and header", `<header class="wp-content">`, true},

		// Negative test
		{"Only header tag", `<header class="header">`, false},
		{"Only wp-content", `<div class="wp-content"></div>`, false},
		{"No relevant content", `<div>Content</div>`, false},
		{"Empty string", ``, false},
		// Edge
		{"Malformed but contains both", `<header><div class="wp-content"`, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := wordpress.MatchHeader(tt.html)
			if result != tt.expected {
				t.Errorf("MatchHeader() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestExtractHeader(t *testing.T) {
	tests := []struct {
		name       string
		html       string
		expectErr  bool
		expectFind bool
	}{
		// Positive
		{
			name:       "Valid header with wp-content",
			html:       `<header class="wp-content">Header</header>`,
			expectErr:  false,
			expectFind: true,
		},

		// Negative
		{
			name:       "No header present",
			html:       `<div class="wp-content">No header</div>`,
			expectErr:  true,
			expectFind: false,
		},
		{
			name:       "Empty HTML input",
			html:       ``,
			expectErr:  true,
			expectFind: false,
		},

		// Edge
		{
			name:       "Malformed HTML but header exists",
			html:       `<header class="wp-content">Unclosed`,
			expectErr:  false,
			expectFind: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block, err := wordpress.ExtractHeader(tt.html, "http://test.site")

			if (err != nil) != tt.expectErr {
				t.Errorf("Expected error = %v, got = %v", tt.expectErr, err)
			}
			if (block.Found == "true") != tt.expectFind {
				t.Errorf("Expected Found = %v, got = %s", tt.expectFind, block.Found)
			}
			if block.Type != "" && block.Type != "wordpress_header" {
				t.Errorf("Unexpected Type = %s", block.Type)
			}
			if block.PageURL != "http://test.site" {
				t.Errorf("PageURL mismatch")
			}
		})
	}
}
