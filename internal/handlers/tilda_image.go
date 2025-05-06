package handlers

import (
	"web-crawler/internal/model"
	"web-crawler/internal/parser/tilda"
)

type TildaImgHandler struct{}

func (t *TildaImgHandler) Match(html string) bool {
	return tilda.MatchImage(html)
}

func (t *TildaImgHandler) Extract(html, pageURL string) (model.Block, error) {
	return tilda.ExtractImage(html, pageURL)
}

func (t *TildaImgHandler) Type() string {
	return "tilda_image"
}

func init() {
	model.RegisterHandler(&TildaImgHandler{})
}
