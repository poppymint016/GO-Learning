package main

import (
	"GO-Project/configs"
	"GO-Project/handlers"
	"GO-Project/services"
	"context"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db := configs.ConnectDB()
	defer func() {
		if err := db.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	expService := services.NewExperienceService(db)
	handlers.NewExperienceHandle(app, expService)

	if err := app.Listen(":6000"); err != nil {
		panic(err)
	}
}
