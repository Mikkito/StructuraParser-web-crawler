package html5

import (
	"errors"
	"strings"
	"web-crawler/internal/model"

	"github.com/PuerkitoBio/goquery"
)

func MatchImage(html string) bool {
	return strings.Contains(html, "<img")
}

func ExtractImage(html, pageURL string) (model.Block, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return model.Block{}, err
	}

	img := doc.Find("img")
	if img.Length() == 0 {
		return model.Block{
			Type:     "html5_img",
			HTML:     "",
			PageURL:  pageURL,
			Found:    "false",
			Accuracy: "0.0",
		}, errors.New("Image not found")
	}

	htmlBlock, err := goquery.OuterHtml(img.First())
	if err != nil {
		return model.Block{}, err
	}

	return model.Block{
		Type:     "html5_img",
		HTML:     htmlBlock,
		PageURL:  pageURL,
		Found:    "true",
		Accuracy: "0.7",
	}, nil
}
