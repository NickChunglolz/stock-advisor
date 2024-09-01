package main

import (
	"log"
	"os"

	"github.com/NickChunglolz/stock-advisor/ochestractor-dms/adapter/controller"
	"github.com/NickChunglolz/stock-advisor/ochestractor-dms/infrastructure/client"
	"github.com/NickChunglolz/stock-advisor/ochestractor-dms/infrastructure/repository"
	utils "github.com/NickChunglolz/stock-advisor/ochestractor-dms/infrastructure/utils"
	"github.com/NickChunglolz/stock-advisor/ochestractor-dms/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

const (
	HOST = "HOST"
	PORT = "PORT"
)

func main() {
	app := fiber.New()

	godotenv.Load()
	address := os.Getenv(HOST) + ":" + os.Getenv(PORT)

	var factory utils.DatabaseFactoryInterface = utils.NewDatabaseFactory()
	defer factory.CloseDatabaseConnections()

	db, err := factory.CreatePostgres()
	if err != nil {
		log.Fatal(err)
	}

	_, err = factory.CreateRedis()
	if err != nil {
		log.Fatal(err)
	}

	clients := client.InitClients()
	repos := repository.InitRepositories(db)
	usecases := usecase.InitUsecases(clients, repos)
	controllers := controller.InitControllers(usecases)
	controller.SetupRoutes(app, controllers)

	log.Fatal(app.Listen(address))
}
