package handlers

import (
	"web-crawler/internal/model"
	"web-crawler/internal/parser/wordpress"
)

type WordPressFooterHandler struct{}

func (h *WordPressFooterHandler) Match(html string) bool {
	return wordpress.MatchHeader(html)
}

func (h *WordPressFooterHandler) Extract(html string, pageURL string) (model.Block, error) {
	return wordpress.ExtractFooter(html, pageURL)
}

func (h *WordPressFooterHandler) Type() string {
	return "wordpress_footer"
}

func init() {
	model.RegisterHandler(&WordPressFooterHandler{})
}
