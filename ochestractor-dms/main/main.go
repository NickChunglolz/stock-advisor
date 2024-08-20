package main

import (
	"log"
	"os"

	"github.com/NickChunglolz/stock-advisor/ochestractor-dms/adapter/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

const (
	host = "HOST"
	port = "PORT"
)

func main() {
	app := fiber.New()
	controller.SetupRoutes(app)

	godotenv.Load()
	address := os.Getenv(host) + ":" + os.Getenv(port)

	log.Fatal(app.Listen(address))

	// factory := &utils.DatabaseFactory{}
	// defer factory.CloseDatabaseConnections()

	// // Initialize Redis client
	// redisClient := factory.CreateRedis()

	// fmt.Println(redisClient)
}
