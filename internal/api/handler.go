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

// Crawler start
func StartCrawlHandler(w http.ResponseWriter, r *http.Request) {
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
	var data map[string]string
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}
	mutx.Lock()
	defer mutx.Unlock()
	crawlerQueue = crawler.NewURLQueue(1000, 2*time.Second)
	crawlerQueue.Enqueue("https://botcreators.ru")
	crawlerQueue.Enqueue("https://structura.app")
	crawlerQueue.Enqueue("https://automatisation.art")
	for _, site := range data {
		crawlerQueue.Enqueue(site)
	}
	go crawlerQueue.Worker()
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
