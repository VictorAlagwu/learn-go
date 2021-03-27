package search

import (
	"encoding/json"
	"os"
)

type Feed struct {
	Name string `json:"site"`
	URI string `json:"link"`
	Type string `json:"type"`
}

const dataFile = "data/data.json"

func RetrieveFeeds() ([]*Feed, error) {
	var feeds []*Feed
	file, err := os.Open(dataFile)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err
}