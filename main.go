package main

import (
	"log"
	"net/http"

	"github.com/kmikko/placeholderize/image"
	"github.com/kmikko/placeholderize/request"
	"github.com/kmikko/placeholderize/response"
)

const ADDR = ":8888"

func handlePNG(w http.ResponseWriter, r *http.Request) {
	log.Printf("New PNG request: %s", r.URL.Query())
	imageRequest := request.NewImageRequest(r)

	img, err := image.CreateImage(imageRequest.Width, imageRequest.Height, imageRequest.Text, imageRequest.Size)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response.WriteResponse(w, img, "png")
}

func handleJPG(w http.ResponseWriter, r *http.Request) {
	log.Printf("New JPG request: %s", r.URL.Query())
	imageRequest := request.NewImageRequest(r)

	img, err := image.CreateImage(imageRequest.Width, imageRequest.Height, imageRequest.Text, imageRequest.Size)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response.WriteResponse(w, img, "jpg")
}

func main() {
	http.HandleFunc("/image.png", handlePNG)
	http.HandleFunc("/image.jpg", handleJPG)
	http.HandleFunc("/image.jpeg", handleJPG)

	log.Printf("Listening on %s", ADDR)
	err := http.ListenAndServe(ADDR, nil)
	if err != nil {
		log.Panicf("Could not start: %s\n", err)
	}
}
