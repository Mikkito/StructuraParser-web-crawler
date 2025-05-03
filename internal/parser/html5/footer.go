package html5

import (
	"errors"
	"strings"
	"web-crawler/internal/model"

	"github.com/PuerkitoBio/goquery"
)

func MatchFooter(html string) bool {
	isNotWordPress := !strings.Contains(html, "wp-content")
	isNotTilda := !strings.Contains(html, "tilda-block")
	isNotBitrix := !strings.Contains(html, "bitrix")
	return strings.Contains(html, "<footer") && isNotBitrix && isNotTilda && isNotWordPress
}

func ExtractFooter(html, pageURL string) (model.Block, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return model.Block{}, err
	}
	footer := doc.Find("footer")
	if footer.Length() == 0 {
		return model.Block{
			Type:     "html5_footer",
			HTML:     "",
			PageURL:  pageURL,
			Found:    "false",
			Accuracy: "0.0",
		}, errors.New("footer not found")
	}
	htmlBlock, err := goquery.OuterHtml(footer.First())
	if err != nil {
		return model.Block{}, nil
	}
	return model.Block{
		Type:     "html5_footer",
		HTML:     htmlBlock,
		PageURL:  pageURL,
		Found:    "true",
		Accuracy: "0.85",
	}, nil
}
