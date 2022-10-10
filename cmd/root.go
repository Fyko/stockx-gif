package cmd

import (
	"fmt"
	"os"

	"github.com/Fyko/stockx-gif/pkg/router"

	stxlogger "github.com/Fyko/stockx-gif/pkg/logger"
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/cobra"
)

var log = stxlogger.CreateLogger()

var port = os.Getenv("PORT")

var rootCmd = &cobra.Command{
	Use:   "stockx-gif",
	Short: "Generate 360° GIFs from StockX listings",
	Long: `StockX GIF Generator is a server and CLI tool that generates 360° GIFs from StockX listings.
				  Source code at http://github.com/Fyko/stockx-gif#README`,
	Run: func(cmd *cobra.Command, args []string) {
		app := fiber.New(fiber.Config{
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return c.Status(500).SendString(err.Error())
			},
			Immutable: true,
		})
		app.Use(logger.New(logger.Config{
			TimeZone: cmd.Flag("timezone").Value.String(),
		}))
		prometheus := fiberprometheus.NewWith("stockx_gif_api", "stockx_gif", "http")
		prometheus.RegisterAt(app, "/api/metrics")
		app.Use(prometheus.Middleware)

		router.SetupRoutes(app)

		port = func() string {
			if len(port) > 0 {
				return port
			} else {
				return cmd.Flags().Lookup("port").Value.String()
			}
		}()
		log.Fatal(app.Listen(fmt.Sprintf(":%v", port)))
	},
}

func init() {
	rootCmd.Flags().Int("port", 2319, "Port to run the server on")
	rootCmd.Flags().String("timezone", "America/Denver", "Timezone to use for logging")
}
