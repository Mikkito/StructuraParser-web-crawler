package dispatcher

import (
	"fmt"
	"web-crawler/internal/model"
)

func Dispatch(html string, pageURL string) ([]model.Block, error) {
	var blocks []model.Block
	for _, handler := range model.GetAllHandlers() {
		if handler.Match(html) {
			block, err := handler.Extract(html, pageURL)
			if err == nil {
				blocks = append(blocks, block)
			}
		}
	}
	if len(blocks) == 0 {
		return nil, fmt.Errorf("no suitable blocks found")
	}
	return blocks, nil
}

// for test
func ResetHandlers() {
	handlers := model.GetAllHandlers()
	for k := range handlers {
		delete(handlers, k)
	}
}
