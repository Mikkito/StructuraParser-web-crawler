package handlers

import (
	"web-crawler/internal/model"
	"web-crawler/internal/parser/tilda"
)

type TildaHeaderHandler struct{}

func (h *TildaHeaderHandler) Match(html string) bool {
	return tilda.MatchHeader(html)
}

func (h *TildaHeaderHandler) Extract(html, pageURL string) (model.Block, error) {
	return tilda.ExtractHeader(html, pageURL)
}
func (h *TildaHeaderHandler) Type() string {
	return "tilda_header"
}
func init() {
	model.RegisterHandler(&TildaHeaderHandler{})
}
