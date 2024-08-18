package main

import (
	"fmt"

	"github.com/NickChunglolz/stock-advisor/ochestractor-dms/infrastructure/utils"
)

func main() {
	// app := fiber.New()
	// apiClient := client.NewTwseApiClient()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	url := "/getStockInfo.jsp?ex_ch=tse_0050.tw|tse_0056.tw|tse_2330.tw|tse_2317.tw|tse_1216.tw|otc_6547.tw|otc_6180.tw"
	// 	data, err := apiClient.GetStockInfos(url)
	// 	if err != nil {
	// 		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	// 	}
	// 	return c.JSON(data)
	// })

	// log.Fatal(app.Listen(":3000"))

	factory := &utils.DatabaseFactory{}
	defer factory.CloseDatabaseConnections()

	// Initialize Redis client
	redisClient := factory.CreateRedis()

	fmt.Println(redisClient)
}
