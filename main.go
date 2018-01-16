package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/math/fixed"
)

var colors = []color.RGBA{
	color.RGBA{187, 59, 86, 255},
	color.RGBA{213, 147, 172, 255},
	color.RGBA{221, 187, 137, 255},
	color.RGBA{208, 195, 186, 255},
	color.RGBA{122, 175, 3, 255}}

var dpi float64 = 72

func main() {
	rand.Seed(time.Now().Unix())

	http.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Query())

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

		// Read the font data.
		f, err := truetype.Parse(goregular.TTF)
		if err != nil {
			log.Println(err)
			return
		}

		// Create image
		img := image.NewRGBA(image.Rect(0, 0, width, height))
		draw.Draw(img, img.Bounds(), &image.Uniform{colors[rand.Intn(len(colors))]}, image.ZP, draw.Src)

		// Draw the text
		h := font.HintingNone
		d := &font.Drawer{
			Dst: img,
			Src: image.White,
			Face: truetype.NewFace(f, &truetype.Options{
				Size:    size,
				DPI:     dpi,
				Hinting: h,
			}),
		}

		d.Dot = fixed.Point26_6{
			X: (fixed.I(width) - d.MeasureString(text)) / 2,
			Y: (fixed.I(height + int(math.Ceil(size*dpi/72)))) / 2,
		}
		d.DrawString(text)

		w.Header().Set("Content-Type", "image/png")
		png.Encode(w, img)
	})

	http.ListenAndServe(":8888", nil)
}
