package crawler

import (

	//"fmt"
	"io"
	//"io/ioutil"
	"net/http"
	"web-crawler/internal/dispatcher"
	"web-crawler/internal/model"
	"web-crawler/pkg/utils/logger"

	"golang.org/x/net/html/charset"
	//"os"
	//"strings"
	//"time"
)

func scrapeURL(url string, resultChan chan<- model.Block) {
	logger := logger.Sugared()
	resp, err := http.Get(url)
	if err != nil {
		logger.Errorf("Error fetching URL: ", url, err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		logger.Errorf("Non OK status code: %d for url: %s\n", resp.StatusCode, url)
		return
	}
	reader, err := charset.NewReader(resp.Body, resp.Header.Get("Content-type"))
	if err != nil {
		logger.Errorf("Encoding error: ", err)
		return
	}
	htmlBytes, err := io.ReadAll(reader)
	if err != nil {
		logger.Errorf("Error read HTML: ", err)
		return
	}
	html := string(htmlBytes)
	block, err := dispatcher.Dispatch(html, url)
	if err != nil {
		logger.Infof("Block not found or error: %v", err)
		return
	}
	resultChan <- block
}
