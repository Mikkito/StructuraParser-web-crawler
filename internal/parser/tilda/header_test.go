package tilda_test

import (
	"testing"
	"web-crawler/internal/parser/tilda"
)

func TestMatchHeader(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected bool
	}{
		// Positive
		{"Contains tilda-block and header", `<header class="tilda-block">`, true},

		// Negative
		{"Only header tag", `<header class="main-header">`, false},
		{"Only tilda-block", `<div class="tilda-block"></div>`, false},
		{"No relevant content", `<div>Just content</div>`, false},
		{"Empty string", ``, false},

		// Edge
		{"Malformed but contains both", `<header><div class="tilda-block"`, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tilda.MatchHeader(tt.html)
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
			name:       "Valid header with tilda-block",
			html:       `<header class="tilda-block">Header</header>`,
			expectErr:  false,
			expectFind: true,
		},

		// Negative
		{
			name:       "No header present",
			html:       `<div class="tilda-block">No header</div>`,
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
			html:       `<header class="tilda-block">Unclosed`,
			expectErr:  false,
			expectFind: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block, err := tilda.ExtractHeader(tt.html, "http://test.site")

			if (err != nil) != tt.expectErr {
				t.Errorf("Expected error = %v, got = %v", tt.expectErr, err)
			}
			if (block.Found == "true") != tt.expectFind {
				t.Errorf("Expected Found = %v, got = %s", tt.expectFind, block.Found)
			}
			if block.Type != "" && block.Type != "tilda_header" {
				t.Errorf("Unexpected Type = %s", block.Type)
			}
			if block.PageURL != "http://test.site" {
				t.Errorf("PageURL mismatch")
			}
		})
	}
}
