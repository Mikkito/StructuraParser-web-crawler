package handlers

import (
	"web-crawler/internal/model"
	"web-crawler/internal/parser/bitrix"
)

type BitrixImgHandler struct{}

func (t *BitrixImgHandler) Match(html string) bool {
	return bitrix.MatchImage(html)
}

func (t *BitrixImgHandler) Extract(html, pageURL string) (model.Block, error) {
	return bitrix.ExtractImage(html, pageURL)
}

func (t *BitrixImgHandler) Type() string {
	return "bitrix_image"
}

func init() {
	model.RegisterHandler(&BitrixImgHandler{})
}
