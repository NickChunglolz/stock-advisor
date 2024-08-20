package usecase

import "github.com/NickChunglolz/stock-advisor/ochestractor-dms/infrastructure/client"

const (
	mockQueryString = "tse_0050.tw|tse_0056.tw|tse_2330.tw|tse_2317.tw|tse_1216.tw|otc_6547.tw|otc_6180.tw"
)

// var mockQueryParams = []string{"tse_0050.tw", "tse_0056.tw", "tse_2330.tw", "tse_2317.tw", "tse_1216.tw", "otc_6547.tw", "otc_6180.tw"}

func GetStocks() ([]client.StockInfoReply, error) {
	apiClient := client.NewTwseApiClient()
	return apiClient.GetStockInfos(mockQueryString)
}