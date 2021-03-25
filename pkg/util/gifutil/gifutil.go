package gifutil

import (
	"bytes"
	"image/gif"
	"strings"

	"github.com/fyko/stockx-gif/next/pkg/logger"
	"github.com/fyko/stockx-gif/next/pkg/util/download"
)

var log = logger.CreateLogger()

// Cretes a gif from the provided urls
func WriteGIF(urls []string) (*bytes.Buffer, error) {
	generator := &gif.GIF{}

	files, err := download.DownloadAndEncodeAll(urls)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(files); i++ {
		frame := files[i]

		generator.Image = append(generator.Image, frame)
		generator.Delay = append(generator.Delay, 0)
		log.Infof("Successfully added image #%v to the generator", i)
	}

	out := new(bytes.Buffer)
	gif.EncodeAll(out, generator)

	return out, nil
}

func Prepare360Images(images []string) []string {
	var mutated []string

	for _, image := range images {
		split := strings.Split(image, "?")
		mutated = append(mutated, split[1]+"?w=1280&fm=jpg")
	}

	return mutated
}
