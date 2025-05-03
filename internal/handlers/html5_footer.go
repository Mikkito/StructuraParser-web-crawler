package handlers

import (
	"web-crawler/internal/model"
	"web-crawler/internal/parser/html5"
)

type Html5FooterHandler struct{}

func (h *Html5FooterHandler) Match(html string) bool {
	return html5.MatchFooter(html)
}

func (h *Html5FooterHandler) Extract(html, pageURL string) (model.Block, error) {
	return html5.ExtractFooter(html, pageURL)
}

func (h *Html5FooterHandler) Type() string {
	return "html5_footer"
}

func init() {
	model.RegisterHandler(&Html5FooterHandler{})
}
