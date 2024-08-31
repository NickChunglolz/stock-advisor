package main

import (
	"log"
	"os"

	"github.com/NickChunglolz/stock-advisor/ochestractor-dms/adapter/controller"
	utils "github.com/NickChunglolz/stock-advisor/ochestractor-dms/infrastructure/utils"
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

	var factory utils.DatabaseFactoryInterface = utils.NewDatabaseFactory()
	defer factory.CloseDatabaseConnections()

	_, err := factory.CreatePostgres()
	if err != nil {
		log.Fatal(err)
	}

	_, err = factory.CreateRedis()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(app.Listen(address))
}
