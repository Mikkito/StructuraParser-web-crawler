package handlers

import (
	"web-crawler/internal/model"
	"web-crawler/internal/parser/wordpress"
)

type WordPressHeaderHandler struct{}

func (h *WordPressHeaderHandler) Match(html string) bool {
	return wordpress.MatchHeader(html)
}

func (h *WordPressHeaderHandler) Extract(html string, pageURL string) (model.Block, error) {
	return wordpress.ExtractHeader(html, pageURL)
}

func (h *WordPressHeaderHandler) Type() string {
	return "wordpress_header"
}

func init() {
	model.RegisterHandler(&WordPressHeaderHandler{})
}
