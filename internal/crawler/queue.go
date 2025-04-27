package crawler

import (
	"sync"
	"time"
)

type URLQueue struct {
	urls  chan string
	wg    sync.WaitGroup
	delay time.Duration
}

func NewURLQueue(size int, delay time.Duration) *URLQueue {
	return &URLQueue{
		urls:  make(chan string, size),
		delay: delay,
	}
}

func (q *URLQueue) Enqueue(url string) {
	q.urls <- url
}

func (q *URLQueue) Dequeue() string {
	return <-q.urls
}

// Starting the processing process

func (q *URLQueue) Worker() {
	defer q.wg.Done()
	for url := range q.urls {
		time.Sleep(q.delay) // delay with request
		go scrapeURL(url)   // processing url
	}
}

func (q *URLQueue) Wait() {
	q.wg.Wait()
}
