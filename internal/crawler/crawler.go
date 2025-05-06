package crawler

import (
	"io"
	"net/http"
	"web-crawler/internal/dispatcher"
	"web-crawler/internal/model"

	"go.uber.org/zap"
	"golang.org/x/net/html/charset"
)

func ScrapeURL(url string, resultChan chan<- model.Block, log *zap.SugaredLogger) {
	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("Error fetching URL: ", url, err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Errorf("Non OK status code: %d for url: %s\n", resp.StatusCode, url)
		return
	}
	reader, err := charset.NewReader(resp.Body, resp.Header.Get("Content-type"))
	if err != nil {
		log.Errorf("Encoding error: ", err)
		return
	}
	htmlBytes, err := io.ReadAll(reader)
	if err != nil {
		log.Errorf("Error read HTML: ", err)
		return
	}
	html := string(htmlBytes)
	block, err := dispatcher.Dispatch(html, url)
	if err != nil {
		log.Infof("Block not found or error: %v", err)
		return
	}
	resultChan <- block
}
