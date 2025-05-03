package tilda

import (
	"errors"
	"strings"
	"web-crawler/internal/model"

	"github.com/PuerkitoBio/goquery"
)

func MatchFooter(html string) bool {
	return strings.Contains(html, "tilda-block") && strings.Contains(html, "<footer")
}

func ExtractFooter(html, pageURL string) (model.Block, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return model.Block{}, err
	}
	footer := doc.Find("footer")
	if footer.Length() == 0 {
		return model.Block{
			Type:     "tilda_footer",
			HTML:     "",
			PageURL:  pageURL,
			Found:    "false",
			Accuracy: "0.0",
		}, errors.New("header not found")
	}
	htmlBlock, err := goquery.OuterHtml(footer.First())
	if err != nil {
		return model.Block{}, err
	}
	return model.Block{
		Type:     "tilda_footer",
		HTML:     htmlBlock,
		PageURL:  pageURL,
		Found:    "true",
		Accuracy: "0.85",
	}, nil
}
