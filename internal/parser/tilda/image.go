package tilda

import (
	"errors"
	"strings"
	"web-crawler/internal/model"

	"github.com/PuerkitoBio/goquery"
)

func MatchImage(html string) bool {
	return strings.Contains(html, `class="t-img"`) && strings.Contains(html, "<img")
}

func ExtractImage(html, pageURL string) (model.Block, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return model.Block{}, err
	}
	img := doc.Find("div.t-img img").First()
	if img.Length() == 0 {
		return model.Block{}, errors.New("Img not found")
	}
	htmlImg, err := goquery.OuterHtml(img)
	if err != nil {
		return model.Block{}, err
	}
	return model.Block{
		Type:     "tilda_img",
		HTML:     htmlImg,
		PageURL:  pageURL,
		Found:    "true",
		Accuracy: "0.75",
	}, nil

}
