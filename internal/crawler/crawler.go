package crawler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
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
	// Save data for ml DON`T FORGET DELETE IN PRODACTION VERSION !!!!
	dirPath := "web-crawler/results"
	err = CreateDirIfNotExist(dirPath)
	if err != nil {
		fmt.Printf("Error dir created: %v", err)
		return
	}
	fileName := fmt.Sprintf("%s_%d.html", sanitizeURL(url), time.Now().Unix())
	err = SaveToFile(dirPath, fileName, string(htmlBytes))
	if err != nil {
		fmt.Printf("Save error: %v", err)
	}
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

// For ML don`t forget delete on prodaction
func CreateDirIfNotExist(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return os.MkdirAll(dirPath, os.ModePerm)
	}
	return nil
}
func SaveToFile(dirPath, fileName, content string) error {
	filePath := fmt.Sprintf("%s/%s", dirPath, fileName)
	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("Error writing to file: %v", err)
	}
	fmt.Printf("The file was saved successfully %s\n", filePath)
	return nil
}
func sanitizeURL(url string) string {
	return strings.ReplaceAll(strings.ReplaceAll(url, "https://", ""), "/", "_")
}
