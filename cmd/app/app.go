package main

import (
	"web-crawler/cmd/server"
	_ "web-crawler/internal/handlers"
)

func main() {
	server.StartServer()
}
