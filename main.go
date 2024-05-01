package main

import (
	"fmt"
	route "go-filestore/api/routes"
	config "go-filestore/configs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	db := config.NewDBConnection()
	app.Static("/buckets", "./buckets")

	route.Setup(app, db)

	err := app.Listen(fmt.Sprintf(":%s", config.GetEnv("app.port", "3000")))
	if err != nil {
		panic(err.Error())
	}
}
