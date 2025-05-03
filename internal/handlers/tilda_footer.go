package handlers

import (
	"web-crawler/internal/model"
	"web-crawler/internal/parser/tilda"
)

type TildaFooterHandler struct{}

func (h *TildaFooterHandler) Match(html string) bool {
	return tilda.MatchFooter(html)
}

func (h *TildaFooterHandler) Extract(html, pageURL string) (model.Block, error) {
	return tilda.ExtractFooter(html, pageURL)
}

func (h *TildaFooterHandler) Type() string {
	return "tilda_footer"
}

func init() {
	model.RegisterHandler(&TildaFooterHandler{})
}
