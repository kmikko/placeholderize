package response

import (
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, img image.Image, t string) {
	switch t {
	case "png":
		w.Header().Set("Content-Type", "image/png")
		png.Encode(w, img)
	case "jpg", "jpeg":
		jpeg.Encode(w, img, &jpeg.Options{Quality: 80})
	}

}
