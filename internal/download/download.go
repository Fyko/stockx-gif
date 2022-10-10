package download

import (
	"bytes"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"io"
	"net/http"

	"github.com/Fyko/stockx-gif/pkg/logger"
)

var log = logger.CreateLogger()

// Does all the encoding magic on a file
func PrepareFile(i int, data io.ReadCloser) (*image.Paletted, error) {
	decoded, err := jpeg.Decode(data)
	if err != nil {
		return nil, err
	}
	log.Infof("Successfully decoded #%v from JPEG", i)

	encoded := new(bytes.Buffer)
	err = gif.Encode(encoded, decoded, nil)
	if err != nil {
		return nil, err
	}
	log.Infof("Successfully encoded #%v to GIF", i)

	frame, err := gif.Decode(bytes.NewReader(encoded.Bytes()))
	if err != nil {
		return nil, err
	}
	log.Infof("Successfully decoded #%v to GIF", i)

	return frame.(*image.Paletted), nil
}

func download(i int, url string) (*image.Paletted, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	log.Infof("Successfully downloaded image #%v", i)

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}

	prep, err := PrepareFile(i, response.Body)
	if err != nil {
		return nil, err
	}

	return prep, nil
}

// Downloads many files and prepares them for gif shit
func DownloadAndEncodeAll(urls []string) (map[int]*image.Paletted, error) {
	files := make(map[int]*image.Paletted)
	done := make(chan *image.Paletted, len(urls))
	errch := make(chan error, len(urls))

	for i, url := range urls {
		go func(i int, url string) {
			b, err := download(i, url)
			if err != nil {
				errch <- err
				done <- nil
				return
			}
			files[i] = b

			done <- b
			errch <- nil
		}(i, url)
	}

	var errStr string
	for i := 0; i < len(urls); i++ {
		if err := <-errch; err != nil {
			errStr = errStr + " " + err.Error()
		}
	}

	var err error
	if errStr != "" {
		err = errors.New(errStr)
	}

	return files, err
}
