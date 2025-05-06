package tilda_test

import (
	"testing"
	"web-crawler/internal/parser/tilda"
)

func TestMatchFooter(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected bool
	}{
		// Positive
		{"Contains tilda-block and footer", `<footer class="tilda-block">`, true},

		// Negative
		{"Only footer tag", `<footer class="main-footer">`, false},
		{"Only tilda-block", `<div class="tilda-block"></div>`, false},
		{"No relevant content", `<div>Just content</div>`, false},
		{"Empty string", ``, false},

		// Edge
		{"Malformed but contains both", `<footer><div class="tilda-block"`, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tilda.MatchFooter(tt.html)
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
			name:       "Valid footer with tilda-block",
			html:       `<footer class="tilda-block">Footer</footer>`,
			expectErr:  false,
			expectFind: true,
		},

		// Negative
		{
			name:       "No footer present",
			html:       `<div class="tilda-block">No footer</div>`,
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
			html:       `<footer class="tilda-block">Unclosed`,
			expectErr:  false,
			expectFind: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block, err := tilda.ExtractFooter(tt.html, "http://test.site")

			if (err != nil) != tt.expectErr {
				t.Errorf("Expected error = %v, got = %v", tt.expectErr, err)
			}
			if (block.Found == "true") != tt.expectFind {
				t.Errorf("Expected Found = %v, got = %s", tt.expectFind, block.Found)
			}
			if block.Type != "" && block.Type != "tilda_footer" {
				t.Errorf("Unexpected Type = %s", block.Type)
			}
			if block.PageURL != "http://test.site" {
				t.Errorf("PageURL mismatch")
			}
		})
	}
}
