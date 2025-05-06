package handlers

import (
	"web-crawler/internal/model"
	"web-crawler/internal/parser/wordpress"
)

type WordPressImgHandler struct{}

func (t *WordPressImgHandler) Match(html string) bool {
	return wordpress.MatchImage(html)
}

func (t *WordPressImgHandler) Extract(html, pageURL string) (model.Block, error) {
	return wordpress.ExtractImage(html, pageURL)
}

func (t *WordPressImgHandler) Type() string {
	return "wordpress_image"
}

func init() {
	model.RegisterHandler(&WordPressImgHandler{})
}
