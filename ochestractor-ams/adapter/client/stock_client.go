package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type StockReply struct {
	Exchange      string `json:"ex"`
	StockCode     string `json:"ch"`
	StockName     string `json:"n"`
	CurrentPrice  string `json:"z"`
	OpeningPrice  string `json:"o"`
	HighestPrice  string `json:"h"`
	LowestPrice   string `json:"l"`
	PreviousClose string `json:"y"`
	TradingVolume string `json:"tv"`
}

type StockClient struct {
	baseURL    string
	httpClient *http.Client
}

const (
	downstreamHost = "DOWNSTREAM_HOST"
	downstreamPort = "DOWNSTREAM_PORT"
	domainPath     = "Stocks"
)

func NewStockClient() *StockClient {
	return &StockClient{
		baseURL:    "http://" + os.Getenv(downstreamHost) + ":" + os.Getenv(downstreamPort) + "/" + domainPath,
		httpClient: &http.Client{},
	}
}

func (client *StockClient) GetStocks() ([]StockReply, error) {

	resp, err := client.httpClient.Get(client.baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	var apiResponse []StockReply
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %w", err)
	}

	return apiResponse, nil
}
