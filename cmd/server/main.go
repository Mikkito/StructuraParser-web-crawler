package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
	"web-crawler/internal/crawler"
)

var crawlerQueue *crawler.URLQueue
var mux sync.Mutex

// Crawler start
func startCrawlHandler(w http.ResponseWriter, r *http.Request) {
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
	mux.Lock()
	defer mux.Unlock()
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

func Main() {
	/*err := logger.Init("./internal/utils/logger/config.yaml")
	if err != nil {
		log.Fatalf("Could not initialize logger: %v", err)
	}
	defer logger.Sync()*/
	http.HandleFunc("/start-crawl", startCrawlHandler)
	http.HandleFunc("/status", getStatusHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
