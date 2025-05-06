package wordpress_test

import (
	"testing"
	"web-crawler/internal/parser/wordpress"
)

func TestMatchFooter(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected bool
	}{
		// Positive
		{"Contains wp-content and footer", `<footer class="wp-content">`, true},

		// Negative
		{"Only footer tag", `<footer class="just-footer">`, false},
		{"Only wp-content", `<div class="wp-content"></div>`, false},
		{"No relevant content", `<div>Just content</div>`, false},
		{"Empty string", ``, false},

		// Edge
		{"Malformed but contains both", `<footer><div class="wp-content"`, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := wordpress.MatchFooter(tt.html)
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
			name:       "Valid footer with wp-content",
			html:       `<footer class="wp-content">Footer</footer>`,
			expectErr:  false,
			expectFind: true,
		},

		// Negative
		{
			name:       "No footer present",
			html:       `<div class="wp-content">No footer</div>`,
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
			name:       "Malformed HTML but footer exists",
			html:       `<footer class="wp-content">Unclosed`,
			expectErr:  false,
			expectFind: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block, err := wordpress.ExtractFooter(tt.html, "http://test.site")

			if (err != nil) != tt.expectErr {
				t.Errorf("Expected error = %v, got = %v", tt.expectErr, err)
			}
			if (block.Found == "true") != tt.expectFind {
				t.Errorf("Expected Found = %v, got = %s", tt.expectFind, block.Found)
			}
			if block.Type != "" && block.Type != "wordpress_footer" {
				t.Errorf("Unexpected Type = %s", block.Type)
			}
			if block.PageURL != "http://test.site" {
				t.Errorf("PageURL mismatch")
			}
		})
	}
}
