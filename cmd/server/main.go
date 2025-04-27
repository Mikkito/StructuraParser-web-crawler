package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
	"web-crawler/internal/crawler"
	"web-crawler/internal/utils/logger"
)

var crawlerQueue *crawler.URLQueue
var mux sync.Mutex

// Crawler start
func startCrawlHandler(w http.ResponseWriter, r *http.Request) {
	mux.Lock()
	defer mux.Unlock()
	crawlerQueue = crawler.NewURLQueue(1000, 2*time.Second)
	crawlerQueue.Enqueue("https://botcreators.ru")
	crawlerQueue.Enqueue("https://structura.app")
	crawlerQueue.Enqueue("https://automatisation.art")
	go crawlerQueue.Worker()
	crawlerQueue.Wait()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Crawl started")
}

func getStatusHandler(w http.ResponseWriter, r *http.Request) {
	mux.Lock()
	defer mux.Unlock()

	// Статус краулера
	if crawlerQueue == nil {
		http.Error(w, "Crawl not started", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Crawl in progress")
}

func main() {
	err := logger.Init("./internal/utils/logger/config.yaml")
	if err != nil {
		log.Fatalf("Could not initialize logger: %v", err)
	}
	defer logger.Sync()
	http.HandleFunc("/start-crawl", startCrawlHandler)
	http.HandleFunc("/status", getStatusHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
