package image

import "testing"

func TestCreateImage(t *testing.T) {
	if _, err := CreateImage(100, 100, "Hello world", 50); err != nil {
		t.Error("Could not create image")
	}
}

func TestImageWidth(t *testing.T) {
	img, _ := CreateImage(200, 300, "Hello world", 50)
	if img.Bounds().Dx() != 200 {
		t.Error("Invalid width")
	}
}

func TestImagHeight(t *testing.T) {
	img, _ := CreateImage(200, 300, "Hello world", 50)
	if img.Bounds().Dy() != 300 {
		t.Error("Invalid height")
	}
}
