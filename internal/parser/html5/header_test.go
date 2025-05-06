package html5_test

import (
	"testing"
	"web-crawler/internal/parser/html5"
)

func TestMatchHeader(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected bool
	}{
		// Positive
		{"Contains header and random class", `<header class="random-name">`, true},

		// Negative
		{"Contains header and platform tag", `<header class="wp-content">`, false},
		{"No relevant content", `<div>Just content</div>`, false},
		{"Empty string", ``, false},

		// Edge
		{"Malformed but contains both", `<header><div class="random-name"`, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := html5.MatchHeader(tt.html)
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
			name:       "Valid header with random class name",
			html:       `<header class="class-name">Header</header>`,
			expectErr:  false,
			expectFind: true,
		},

		// Negative
		{
			name:       "Empty HTML input",
			html:       ``,
			expectErr:  true,
			expectFind: false,
		},

		// Edge
		{
			name:       "Malformed HTML but header exists",
			html:       `<header class="class-name">Unclosed`,
			expectErr:  false,
			expectFind: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block, err := html5.ExtractHeader(tt.html, "http://test.site")

			if (err != nil) != tt.expectErr {
				t.Errorf("Expected error = %v, got = %v", tt.expectErr, err)
			}
			if (block.Found == "true") != tt.expectFind {
				t.Errorf("Expected Found = %v, got = %s", tt.expectFind, block.Found)
			}
			if block.Type != "" && block.Type != "html5_header" {
				t.Errorf("Unexpected Type = %s", block.Type)
			}
			if block.PageURL != "http://test.site" {
				t.Errorf("PageURL mismatch")
			}
		})
	}
}
