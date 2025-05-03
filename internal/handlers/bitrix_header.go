package handlers

import (
	"web-crawler/internal/model"
	"web-crawler/internal/parser/bitrix"
)

type BitrixHeaderHandler struct{}

func (h *BitrixHeaderHandler) Match(html string) bool {
	return bitrix.MatchHeader(html)
}

func (h *BitrixHeaderHandler) Extract(html, pageURL string) (model.Block, error) {
	return bitrix.ExtractHeader(html, pageURL)
}

func (h *BitrixHeaderHandler) Type() string {
	return "bitrix_header"
}

func init() {
	model.RegisterHandler(&BitrixHeaderHandler{})
}
