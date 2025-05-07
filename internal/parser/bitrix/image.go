package bitrix

import (
	"errors"
	"strings"
	"web-crawler/internal/model"

	"github.com/PuerkitoBio/goquery"
)

func MatchImage(html string) bool {
	return (strings.Contains(html, `class="bx-image"`) || strings.Contains(html, "bitrix-img") || strings.Contains(html, "bitrix-image")) && strings.Contains(html, "<img")
}

func ExtractImage(html, pageURL string) (model.Block, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return model.Block{}, err
	}
	img := doc.Find("img[class*='bx-image'], img[class*='bitrix-img'], img[class*='bitrix-image']").First()
	if img.Length() == 0 {
		return model.Block{
			Type:     "bitrix_img",
			HTML:     "",
			PageURL:  pageURL,
			Found:    "false",
			Accuracy: "0.0",
		}, errors.New("bitrix img not found")
	}
	htmlBlock, err := goquery.OuterHtml(img)
	if err != nil {
		return model.Block{}, err
	}

	return model.Block{
		Type:     "bitrix_img",
		HTML:     htmlBlock,
		PageURL:  pageURL,
		Found:    "true",
		Accuracy: "0.85",
	}, nil
}
