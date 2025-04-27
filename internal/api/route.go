package api

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/crawl", StartCrawlHandler).Methods("POST")
	router.HandleFunc("/status", GetStatusHandler).Methods("GET")
}
