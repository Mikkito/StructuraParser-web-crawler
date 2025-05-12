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
	log.Infof("Scraping URL: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("Error fetching URL %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Errorf("Non-OK status code: %d for URL: %s", resp.StatusCode, url)
		return
	}

	reader, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		log.Errorf("Encoding error for URL %s: %v", url, err)
		return
	}

	htmlBytes, err := io.ReadAll(reader)
	if err != nil {
		log.Errorf("Error reading HTML from %s: %v", url, err)
		return
	}

	html := string(htmlBytes)
	log.Infof("Dispatching HTML from: %s", url)

	blocks, err := dispatcher.Dispatch(html, url)
	if err != nil {
		log.Warnf("Dispatch error for %s: %v", url, err)
		return
	}

	for _, block := range blocks {
		log.Infof("Dispatched block type: %s", block.Type)
		resultChan <- block
	}
}
