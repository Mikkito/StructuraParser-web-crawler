package model

type ParseRequest struct {
	URL    string   `json: "url"`
	Blocks []string `json: "blocks"`
}
