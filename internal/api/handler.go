package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
	"web-crawler/internal/crawler"
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
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
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
	crawlerQueue := crawler.NewURLQueue(10, 2*time.Second)
	worker := crawler.NewWorker(crawlerQueue)
	worker.Start(5)
	crawlerQueue.Enqueue("https://botcreators.ru")
	crawlerQueue.Enqueue("https://structura.app")
	crawlerQueue.Enqueue("https://automatisation.art")
	for _, url := range req.URLs {
		crawlerQueue.Enqueue(url)
	}
	crawlerQueue.Wait()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Crawl started")
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
