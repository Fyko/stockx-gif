package router

import (
	"stockx-gif-next/pkg/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(app *fiber.App) {
	app.Use(cors.New())
	app.Use(compress.New())

	app.Post("/generate", handler.GenerateGIF)
	app.Get("/query_algolia", handler.FetchAlgoliaData)
	app.Get("/query_stockx", handler.FetchStockXData)
	app.Get("/static", handler.ReturnS3URL)
	app.Use("/ping", handler.Ping)
}
