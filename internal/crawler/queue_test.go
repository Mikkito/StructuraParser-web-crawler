package crawler

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	testUrl := "https://testurl.com"
	testUrl2 := "https://twotesturl.com"
	TestQueue := NewURLQueue(10, 2*time.Second)
	TestQueue.Enqueue(testUrl)
	TestQueue.Enqueue(testUrl2)
	assert.Equal(t, testUrl, TestQueue.Dequeue(), "queue is working correctly")
}
