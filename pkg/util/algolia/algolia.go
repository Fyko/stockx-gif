package algolia

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/fyko/stockx-gif/next/pkg/logger"
)

var log = logger.CreateLogger()

var uri = url.URL{
	Scheme:   "https",
	Host:     "xw7sbct9v6-dsn.algolia.net",
	Path:     "1/indexes/products/query",
	RawQuery: "x-algolia-agent=Algolia%20for%20JavaScript%20(4.8.4)%3B%20Browser",
}

type TruncatedHit struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func TruncateResponse(data AlgolaQueryResponse) []TruncatedHit {
	var hits []TruncatedHit
	for _, hit := range data.Hits {
		if len(hit.ID) > 0 {
			hits = append(hits, TruncatedHit{ID: hit.ID, Name: hit.Name})
		}
	}

	return hits
}

func DoAlgoliaQuery(query string) ([]TruncatedHit, error) {
	data := map[string]string{"params": fmt.Sprintf("query=%s&hitsPerPage=50&facets=*", query)}
	body, _ := json.Marshal(data)

	req, _ := http.NewRequest(http.MethodPost, uri.String(), bytes.NewBuffer(body))
	req.Header.Add("x-algolia-api-key", "6b5e76b49705eb9f51a06d3c82f7acee")
	req.Header.Add("x-algolia-application-id", "XW7SBCT9V6")
	req.Header.Add("Content-Type", "application/json")

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Errorf("An error occurred when fetching Algolia: %v", err)
		return nil, err
	}

	var payload = new(AlgolaQueryResponse)
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&payload)

	if err != nil {
		log.Errorf("An error occurred when decoding Algolia data: %v", err)
		return nil, err
	}

	return TruncateResponse(*payload), nil
}
