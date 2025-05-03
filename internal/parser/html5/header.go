package html5

import (
	"errors"
	"strings"
	"web-crawler/internal/model"

	"github.com/PuerkitoBio/goquery"
)

func MatchHeader(html string) bool {
	isNotWordPress := !strings.Contains(html, "wp-content")
	isNotTilda := !strings.Contains(html, "tilda-block")
	isNotBitrix := !strings.Contains(html, "bitrix")
	return strings.Contains(html, "<header") && isNotBitrix && isNotTilda && isNotWordPress
}

func ExtractHeader(html, pageURL string) (model.Block, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return model.Block{}, err
	}
	header := doc.Find("header")
	if header.Length() == 0 {
		return model.Block{
			Type:     "html5_header",
			HTML:     "",
			PageURL:  pageURL,
			Found:    "false",
			Accuracy: "0.0",
		}, errors.New("header not found")
	}
	htmlBlock, err := goquery.OuterHtml(header.First())
	if err != nil {
		return model.Block{}, nil
	}
	return model.Block{
		Type:     "html5_header",
		HTML:     htmlBlock,
		PageURL:  pageURL,
		Found:    "true",
		Accuracy: "0.85",
	}, nil
}
