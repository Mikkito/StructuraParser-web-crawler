package handlers

import (
	"web-crawler/internal/model"
	"web-crawler/internal/parser/html5"
)

type Html5ImageHandler struct{}

func (h *Html5ImageHandler) Match(html string) bool {
	return html5.MatchImage(html)
}

func (h *Html5ImageHandler) Extract(html, pageURL string) (model.Block, error) {
	return html5.ExtractImage(html, pageURL)
}

func (h *Html5ImageHandler) Type() string {
	return "html5_image"
}

func init() {
	model.RegisterHandler(&Html5ImageHandler{})
}
