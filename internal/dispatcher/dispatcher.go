package dispatcher

import (
	"fmt"
	"web-crawler/internal/model"
)

func Dispatch(html string, pageURL string) (model.Block, error) {
	for _, handler := range model.GetAllHandlers() {
		if handler.Match(html) {
			return handler.Extract(html, pageURL)
		}
	}
	return model.Block{
		Type:     "unknown",
		HTML:     "",
		PageURL:  pageURL,
		Found:    "false",
		Accuracy: "0.0",
	}, fmt.Errorf("no suitable block found")
}
