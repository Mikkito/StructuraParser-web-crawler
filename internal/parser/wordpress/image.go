package wordpress

import (
	"errors"
	"strings"
	"web-crawler/internal/model"

	"github.com/PuerkitoBio/goquery"
)

func MatchImage(html string) bool {
	return strings.Contains(html, "wp-image") || strings.Contains(html, "wp-post-image") && strings.Contains(html, "<img")
}

func ExtractImage(html, pageURL string) (model.Block, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return model.Block{}, err
	}

	img := doc.Find("img[class*='wp-image'], img[class*='wp-post-image']").First()
	if img.Length() == 0 {
		return model.Block{
			Type:     "wordpress_img",
			HTML:     "",
			PageURL:  pageURL,
			Found:    "false",
			Accuracy: "0.0",
		}, errors.New("wordpress img not found")
	}

	htmlBlock, err := goquery.OuterHtml(img)
	if err != nil {
		return model.Block{}, err
	}

	return model.Block{
		Type:     "wordpress_img",
		HTML:     htmlBlock,
		PageURL:  pageURL,
		Found:    "true",
		Accuracy: "0.85",
	}, nil
}
