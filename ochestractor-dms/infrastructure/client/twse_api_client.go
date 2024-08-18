package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type StockInfoReply struct {
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

type StockInfoReplies struct {
	MsgArray []StockInfoReply `json:"msgArray"`
}

type TwseApiClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewTwseApiClient() *TwseApiClient {
	return &TwseApiClient{
		baseURL:    "http://mis.twse.com.tw/stock/api",
		httpClient: &http.Client{},
	}
}

func (client *TwseApiClient) GetStockInfos(url string) ([]StockInfoReply, error) {
	resp, err := client.httpClient.Get(client.baseURL + url)
	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	var apiResponse StockInfoReplies
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %w", err)
	}

	return apiResponse.MsgArray, nil
}
