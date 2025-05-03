package handlers

import (
	"web-crawler/internal/model"
	"web-crawler/internal/parser/html5"
)

type Html5HeaderHandler struct{}

func (h *Html5HeaderHandler) Match(html string) bool {
	return html5.MatchHeader(html)
}

func (h *Html5HeaderHandler) Extract(html, pageURL string) (model.Block, error) {
	return html5.ExtractHeader(html, pageURL)
}

func (h *Html5HeaderHandler) Type() string {
	return "html5_header"
}

func init() {
	model.RegisterHandler(&Html5HeaderHandler{})
}
