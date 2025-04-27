package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"
	"web-crawler/cmd/server"

	"github.com/stretchr/testify/assert"
)

func TestAPI(t *testing.T) {
	go func() {
		server.Main()
	}()
	time.Sleep(5 * time.Second)
	reqBody := map[string]interface{}{
		"1": "https://github.com",
		"2": "https://habr.com",
		"3": "https://dns-shop.ru",
		"4": "https://banki.ru",
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	req, err := http.NewRequest("POST", "http://127.0.0.1:8080/crawl", bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("Request error: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Request send error: %v", err)
	}
	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Ожидался статус 200 OK")
	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Ошибка при декодировании ответа: %v", err)
	}
}
