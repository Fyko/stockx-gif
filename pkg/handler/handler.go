package handler

import (
	"stockx-gif-next/internal/algolia"
	"stockx-gif-next/internal/gifutil"
	"stockx-gif-next/internal/stockx"

	"github.com/gofiber/fiber/v2"
)

type GenerateGifOptions struct {
	ID      string `json:"id"`
	Preview bool   `json:"preview"`
}

type FetchAlgoliaDataOptions struct {
	Query string `query:"query"`
}

type FetchStockxDataOptions struct {
	Product string `query:"product"`
}

func Ping(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).Send([]byte("Hello!"))
}

func GenerateGIF(c *fiber.Ctx) error {
	options := new(GenerateGifOptions)
	if err := c.BodyParser(options); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	product, err := stockx.FetchStockXProductData(options.ID)
	if err != nil {
		return err
	}

	if len(product.Product.Media.The360) == 0 {
		return c.SendStatus(409)
	}

	images := gifutil.Prepare360Images(product.Product.Media.The360, options.Preview)

	gif, err := gifutil.WriteGIF(images)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).Type("gif").SendStream(gif, gif.Len())
}

func FetchAlgoliaData(c *fiber.Ctx) error {
	options := new(FetchAlgoliaDataOptions)
	if err := c.QueryParser(options); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	query, err := algolia.DoAlgoliaQuery(options.Query)
	if err != nil {
		return err
	}

	return c.JSON(query)
}

func FetchStockXData(c *fiber.Ctx) error {
	options := new(FetchStockxDataOptions)
	if err := c.QueryParser(options); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	query, err := stockx.FetchStockXProductData(options.Product)
	if err != nil {
		return err
	}

	return c.JSON(query.Product)
}
