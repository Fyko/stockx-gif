package main

import (
	"fmt"
	"log"
	"os"

	"stockx-gif-next/pkg/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var port = os.Getenv("PORT")

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(500).SendString(err.Error())
		},
		Immutable: true,
	})
	app.Use(logger.New(logger.Config{
		TimeZone: "America/Denver",
	}))

	router.SetupRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%v", port)))
}
