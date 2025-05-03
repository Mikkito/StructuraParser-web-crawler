package model

type Block struct {
	Type     string `json:"type"`
	HTML     string `json:"html"`
	PageURL  string `json:"page_url"`
	Found    string `json:"found"`
	Accuracy string `json:"accuracy"`
}

type BlockHandler interface {
	Match(html string) bool
	Extract(html string, pageURL string) (Block, error)
	Type() string
}
