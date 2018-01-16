package response

import (
	"image"
	"net/http/httptest"
	"testing"
)

func TestPNGContentType(t *testing.T) {
	w := httptest.NewRecorder()
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	WriteResponse(w, img, "png")
	if w.Header().Get("Content-Type") != "image/png" {
		t.Error("Invalid Content-Type")
	}
}

func TestJPGContentType(t *testing.T) {
	w := httptest.NewRecorder()
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	WriteResponse(w, img, "jpeg")
	if w.Header().Get("Content-Type") != "image/jpeg" {
		t.Error("Invalid Content-Type")
	}
}

func TestJPEGContentType(t *testing.T) {
	w := httptest.NewRecorder()
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	WriteResponse(w, img, "jpeg")
	if w.Header().Get("Content-Type") != "image/jpeg" {
		t.Error("Invalid Content-Type")
	}
}
