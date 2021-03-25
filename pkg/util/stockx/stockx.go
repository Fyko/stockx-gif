package stockx

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fyko/stockx-gif/next/pkg/logger"
)

var log = logger.CreateLogger()

func FetchStockXProductData(id string) (*StockXProductResponse, error) {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://stockx.com/api/products/%s?includes=market&currency=usd", id), nil)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:86.0) Gecko/20100101 Firefox/86.0")
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Errorf("An error occurred when fetching StockX: %v", err)
		return nil, err
	}

	if res.StatusCode != 200 {
		log.Errorf("The StockX request returned %v", res.Status)
		return nil, err
	}

	var payload *StockXProductResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&payload)

	if err != nil {
		log.Errorf("An error occurred when decoding StockX data: %v", err)
		return nil, err
	}

	return payload, nil
}
