package request

import (
	"net/http/httptest"
	"testing"
)

func TestImageRequestWidth(t *testing.T) {
	ir := NewImageRequest(httptest.NewRequest("GET", "/image.png?width=300&height=400&text=Hello%20worl&size=50", nil))
	if ir.Width != 300 {
		t.Error("Invalid width")
	}
}

func TestImageRequesHeight(t *testing.T) {
	ir := NewImageRequest(httptest.NewRequest("GET", "/image.png?width=300&height=400&text=Hello%20worl&size=50", nil))
	if ir.Height != 400 {
		t.Error("Invalid height")
	}
}

func TestImageRequestText(t *testing.T) {
	ir := NewImageRequest(httptest.NewRequest("GET", "/image.png?width=300&height=400&text=Hello%20world&size=50", nil))
	if ir.Text != "Hello world" {
		t.Error("Invalid text")
	}
}

func TestImageRequestTextSize(t *testing.T) {
	ir := NewImageRequest(httptest.NewRequest("GET", "/image.png?width=300&height=400&text=Hello%20world&size=50", nil))
	if ir.Size != 50 {
		t.Error("Invalid text size")
	}
}
