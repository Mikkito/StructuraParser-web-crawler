package tilda

import (
	"errors"
	"strings"
	"web-crawler/internal/model"

	"github.com/PuerkitoBio/goquery"
)

func MatchHeader(html string) bool {
	return strings.Contains(html, "tilda-block") && strings.Contains(html, "<header")
}

func ExtractHeader(html, pageURL string) (model.Block, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return model.Block{}, nil
	}
	header := doc.Find("header")
	if header.Length() == 0 {
		return model.Block{
			Type:     "tilda_header",
			HTML:     "",
			PageURL:  pageURL,
			Found:    "false",
			Accuracy: "0.0",
		}, errors.New("header not found")
	}
	htmlBlock, err := goquery.OuterHtml(header.First())
	if err != nil {
		return model.Block{}, err
	}
	return model.Block{
		Type:     "tilda_header",
		HTML:     htmlBlock,
		PageURL:  pageURL,
		Found:    "true",
		Accuracy: "0.85",
	}, nil
}
