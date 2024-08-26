package main

import (
	"fmt"
	"log"
	"os"

	"github.com/NickChunglolz/stock-advisor/ochestractor-dms/adapter/controller"
	"github.com/NickChunglolz/stock-advisor/ochestractor-dms/infrastructure/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

const (
	HOST = "HOST"
	PORT = "PORT"
)

func main() {
	app := fiber.New()
	controller.SetupRoutes(app)

	godotenv.Load()
	address := os.Getenv(HOST) + ":" + os.Getenv(PORT)

	factory := &utils.DatabaseFactory{}
	defer factory.CloseDatabaseConnections()

	postgresClient := factory.CreatePostgres()
	redisClient := factory.CreateRedis()

	fmt.Println(postgresClient)
	fmt.Println(redisClient)

	log.Fatal(app.Listen(address))
}
