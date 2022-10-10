package handler

import (
	"net/http"

	"github.com/Fyko/stockx-gif/internal/algolia"
	"github.com/Fyko/stockx-gif/internal/gifutil"
	"github.com/Fyko/stockx-gif/internal/s3"
	"github.com/Fyko/stockx-gif/internal/stockx"

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

type ReturnS3URLQueryOptions struct {
	ID      string `json:"id"`
	Preview bool   `json:"preview"`
	Fail    bool   `json:"fail"`
}

func Ping(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).Send([]byte("Hello!"))
}

func GenerateGIF(c *fiber.Ctx) error {
	options := new(GenerateGifOptions)
	if err := c.BodyParser(options); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	exists := s3.FetchExisting(options.ID, options.Preview)
	if len(exists) != 0 {
		res, err := http.Get(exists)
		if err == nil {
			return c.Status(fiber.StatusOK).Type("gif").SendStream(res.Body, int(res.ContentLength))
		}
	}

	product, err := stockx.FetchStockXProductData(options.ID)
	if err != nil {
		return err
	}

	if len(product.Product.Media.The360) == 0 {
		return c.SendStatus(409)
	}

	width := func() int {
		if options.Preview {
			return 256
		} else {
			return 1280
		}
	}()

	images := gifutil.Prepare360Images(product.Product.Media.The360, width)

	gif, err := gifutil.WriteGIF(images)
	if err != nil {
		return err
	}

	go func() {
		if len(exists) == 0 {
			s3.UploadGIF(options.ID, options.Preview, gif)
		}
	}()

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

func ReturnS3URL(c *fiber.Ctx) error {
	options := new(ReturnS3URLQueryOptions)
	if err := c.QueryParser(options); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	exists := s3.FetchExisting(options.ID, options.Preview)
	if len(exists) != 0 {
		return c.Status(fiber.StatusOK).SendString(exists)
	} else if options.Fail {
		return c.SendStatus(fiber.StatusNotFound)
	}

	product, err := stockx.FetchStockXProductData(options.ID)
	if err != nil {
		return err
	}

	if len(product.Product.Media.The360) == 0 {
		return c.SendStatus(409)
	}

	width := func() int {
		if options.Preview {
			return 256
		} else {
			return 1280
		}
	}()

	images := gifutil.Prepare360Images(product.Product.Media.The360, width)

	gif, err := gifutil.WriteGIF(images)
	if err != nil {
		return err
	}

	location := s3.UploadGIF(options.ID, options.Preview, gif)

	return c.Status(fiber.StatusOK).SendString(location)
}
