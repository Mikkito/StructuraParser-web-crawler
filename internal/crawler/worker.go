package crawler

import (
	"time"
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
func (w *Worker) Start(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		go w.worker()
	}
}

func (w *Worker) worker() {
	defer w.queue.wg.Done()
	for url := range w.queue.urls {
		time.Sleep(w.queue.delay)
		go scrapeURL(url)
	}
}
