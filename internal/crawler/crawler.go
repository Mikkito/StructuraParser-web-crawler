package crawler

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func scrapeURL(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error fetching URL: ", url, err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Printf("Non OK status code: %d for url: %s\n", resp.StatusCode, url)
		return
	}

	htmlBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error read HTML: ", err)
		return
	}
	// Creating JSON package for Python service
	payload := map[string]string{"html": string(htmlBytes)}
	jsonData, _ := json.Marshal(payload)
	// Send request in pyhton service
	pyResp, err := http.Post("http://localhost:8090/parsehtml", "applications/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error calling Python service: ", err)
		return
	}
	defer pyResp.Body.Close()
	// Result structure
	var result struct {
		Blocks []struct {
			Type     string "json:\"type\""
			Platform string "json:\"platform\""
			Selector string "json:\"selector\""
		} "json:\"blocks\""
	}
	if err := json.NewDecoder(pyResp.Body).Decode(&result); err != nil {
		log.Println("Failed to decode response from ML parser:", err)
		return
	}
	// Working with the result

}
