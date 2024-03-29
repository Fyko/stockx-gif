package gifutil

import (
	"bytes"
	"fmt"
	"image/gif"
	"strings"

	"github.com/Fyko/stockx-gif/internal/download"
	"github.com/Fyko/stockx-gif/pkg/logger"
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

// write docs for this go function
func Prepare360Images(images []string, width int) []string {
	var mutated []string

	for _, image := range images {
		split := strings.Split(image, "?")
		mutated = append(mutated, fmt.Sprintf("%v?w=%v&fm=jpg", split[0], width))
	}

	return mutated
}
