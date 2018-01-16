package request

import (
	"net/http"
	"strconv"
)

type ImageRequest struct {
	Width  int
	Height int
	Size   float64
	Text   string
}

func NewImageRequest(r *http.Request) ImageRequest {
	width, err := strconv.Atoi(r.URL.Query().Get("width"))
	if err != nil {
		width = 640
	}

	height, err := strconv.Atoi(r.URL.Query().Get("height"))
	if err != nil {
		height = 480
	}

	size, err := strconv.ParseFloat(r.URL.Query().Get("size"), 64)
	if err != nil {
		size = 100
	}

	text := r.URL.Query().Get("text")
	if len(text) == 0 {
		text = "Placeholder"
	}

	return ImageRequest{
		Width:  width,
		Height: height,
		Size:   size,
		Text:   text,
	}
}
