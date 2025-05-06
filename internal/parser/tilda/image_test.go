package tilda_test

import (
	"testing"
	"web-crawler/internal/model"
	"web-crawler/internal/parser/tilda"
)

func TestTildaImageParser(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected model.Block
		found    bool
	}{
		{
			name: "Valid Tilda image",
			html: `<html><body><img class="t-img" src="image.jpg"></body></html>`,
			expected: model.Block{
				Type:     "tilda_img",
				Found:    "true",
				HTML:     `<img class="t-img" src="image.jpg">`,
				Accuracy: "0.9",
			},
			found: true,
		},
		{
			name:  "No t-img class",
			html:  `<html><body><img class="other" src="image.jpg"></body></html>`,
			found: false,
		},
		{
			name:  "No img tag",
			html:  `<html><body><div class="t-img"></div></body></html>`,
			found: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if matched := tilda.MatchImage(tt.html); matched != tt.found {
				t.Errorf("MatchImage() = %v, want %v", matched, tt.found)
			}

			if tt.found {
				block, err := tilda.ExtractImage(tt.html, "http://test.site")
				if err != nil {
					t.Fatalf("ExtractImage() unexpected error: %v", err)
				}
				if block.Type != tt.expected.Type || block.Found != tt.expected.Found || block.Accuracy != tt.expected.Accuracy {
					t.Errorf("Extracted block mismatch: %+v", block)
				}
			}
		})
	}
}
