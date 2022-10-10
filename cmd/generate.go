package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Fyko/stockx-gif/internal/algolia"
	"github.com/Fyko/stockx-gif/internal/gifutil"
	"github.com/Fyko/stockx-gif/internal/stockx"
	"github.com/spf13/cobra"
)

var generateCommand = &cobra.Command{
	Use:   "generate <url> --width [1280px]",
	Short: "Generate 360° GIFs from StockX listings",
	Run: func(cmd *cobra.Command, args []string) {
		width := cmd.Flag("width").Value.String()

		x, err := algolia.DoAlgoliaQuery(args[0])
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}
		log.Infof("Found %v results", x[0])

		if len(x) == 0 {
			log.Error("No results found")
			os.Exit(1)
		}

		product, err := stockx.FetchStockXProductData(x[0].ID)
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}

		if len(product.Product.Media.The360) == 0 {
			log.Error("No 360° images found")
			os.Exit(1)
		}

		parsed, err := strconv.ParseInt(width, 10, 0)
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}

		images := gifutil.Prepare360Images(product.Product.Media.The360, int(parsed))

		gif, err := gifutil.WriteGIF(images)
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}

		// write gif (type Buffer) to file
		f, err := os.Create(fmt.Sprintf("%v.gif", x[0].Slug))
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}
		defer f.Close()

		_, err = f.Write(gif.Bytes())
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}

		log.Infof("GIF saved to %s.gif", x[0].Slug)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	generateCommand.Flags().Int32P("width", "w", 1280, "Width of the GIF in pixels (height dynamic)")
	rootCmd.AddCommand(generateCommand)
}
