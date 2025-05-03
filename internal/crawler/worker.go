package crawler

import (
	"sync"
	"time"
	"web-crawler/internal/model"
)

type Worker struct {
	queue *URLQueue
}

func NewWorker(q *URLQueue) *Worker {
	return &Worker{
		queue: q,
	}
}

// Start gorutine pull
func (w *Worker) Start(numWorkers int, resultChan chan<- model.Block, scrapeWg *sync.WaitGroup) {
	for i := 0; i < numWorkers; i++ {
		w.queue.wg.Add(1)
		go w.worker(resultChan, scrapeWg)
	}
}

func (w *Worker) worker(resultChan chan<- model.Block, scrapeWg *sync.WaitGroup) {
	defer w.queue.wg.Done()
	for url := range w.queue.urls {
		time.Sleep(w.queue.delay)
		scrapeWg.Add(1)
		go func(url string) {
			defer scrapeWg.Done()
			scrapeURL(url, resultChan)
		}(url)
	}
}
