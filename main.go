package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/kmikko/placeholderize/image"
	"github.com/kmikko/placeholderize/request"
	"github.com/kmikko/placeholderize/response"
)

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
	var port = flag.Int("port", 8080, "Port on which service will run")
	flag.Parse()

	var ADDR = fmt.Sprintf(":%d", *port)

	http.HandleFunc("/image.png", handlePNG)
	http.HandleFunc("/image.jpg", handleJPG)
	http.HandleFunc("/image.jpeg", handleJPG)

	log.Printf("Listening on %s", ADDR)
	err := http.ListenAndServe(ADDR, nil)
	if err != nil {
		log.Panicf("Could not start: %s\n", err)
	}
}
