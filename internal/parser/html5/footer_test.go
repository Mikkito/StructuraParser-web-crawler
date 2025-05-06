package html5_test

import (
	"testing"
	"web-crawler/internal/parser/html5"
)

func TestMatchFooter(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected bool
	}{
		// Positive
		{"Contains footer and random class", `<footer class="random-name">`, true},

		// Negative
		{"Contains footer and platform tag", `<footer class="wp-content">`, false},
		{"No relevant content", `<div>Just content</div>`, false},
		{"Empty string", ``, false},

		// Edge
		{"Malformed but contains both", `<footer><div class="random-name"`, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := html5.MatchFooter(tt.html)
			if result != tt.expected {
				t.Errorf("MatchFooter() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestExtractFooter(t *testing.T) {
	tests := []struct {
		name       string
		html       string
		expectErr  bool
		expectFind bool
	}{
		// Positive
		{
			name:       "Valid header",
			html:       `<footer class="class-name">Footer</footer>`,
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
			html:       `<footer class="class-name">Unclosed`,
			expectErr:  false,
			expectFind: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block, err := html5.ExtractFooter(tt.html, "http://test.site")

			if (err != nil) != tt.expectErr {
				t.Errorf("Expected error = %v, got = %v", tt.expectErr, err)
			}
			if (block.Found == "true") != tt.expectFind {
				t.Errorf("Expected Found = %v, got = %s", tt.expectFind, block.Found)
			}
			if block.Type != "" && block.Type != "html5_footer" {
				t.Errorf("Unexpected Type = %s", block.Type)
			}
			if block.PageURL != "http://test.site" {
				t.Errorf("PageURL mismatch")
			}
		})
	}
}
