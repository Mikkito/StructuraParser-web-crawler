package handlers

import (
	"web-crawler/internal/model"
	"web-crawler/internal/parser/bitrix"
)

type BitrixFooterHandler struct{}

func (h *BitrixFooterHandler) Match(html string) bool {
	return bitrix.MatchFooter(html)
}

func (h *BitrixFooterHandler) Extract(html, pageURL string) (model.Block, error) {
	return bitrix.ExtractFooter(html, pageURL)
}

func (h *BitrixFooterHandler) Type() string {
	return "bitrix_footer"
}

func init() {
	model.RegisterHandler(&BitrixFooterHandler{})
}
