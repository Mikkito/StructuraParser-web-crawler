package api

import (
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"time"
	"web-crawler/internal/crawler"
	"web-crawler/internal/model"
)

var crawlerQueue *crawler.URLQueue
var mutx sync.Mutex

// Create request structure
type RequestBody struct {
	URLs []string `json:"urls"`
}

// Crawler start
func StartCrawlHandler(w http.ResponseWriter, r *http.Request) {
	var req RequestBody
	var results []model.Block
	var scrapeWg sync.WaitGroup
	resultChan := make(chan model.Block, 100)
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}
	mutx.Lock()
	defer mutx.Unlock()
	//go func() {
	//	for block := range resultChan {
	//		results = append(results, block)
	//	}
	//}()
	crawlerQueue = crawler.NewURLQueue(10, 2*time.Second)
	worker := crawler.NewWorker(crawlerQueue)
	worker.Start(5, resultChan, &scrapeWg)
	for _, url := range req.URLs {
		crawlerQueue.Enqueue(url)
	}
	crawlerQueue.Close()
	crawlerQueue.Wait()
	scrapeWg.Wait()
	close(resultChan)
	for block := range resultChan {
		results = append(results, block)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func GetStatusHandler(w http.ResponseWriter, r *http.Request) {
	mutx.Lock()
	defer mutx.Unlock()

	// Handler status
	if crawlerQueue == nil {
		http.Error(w, "Crawl not started", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Crawl in progress")
}
