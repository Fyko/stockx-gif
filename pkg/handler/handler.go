package handler

import (
	"github.com/fyko/stockx-gif/next/pkg/util/algolia"
	"github.com/fyko/stockx-gif/next/pkg/util/gifutil"
	"github.com/fyko/stockx-gif/next/pkg/util/stockx"
	"github.com/gofiber/fiber/v2"
)

type GenerateGifOptions struct {
	Images []string `json:"images"`
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

	gif, err := gifutil.WriteGIF(options.Images)
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
