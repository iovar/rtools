package tools

import (
	"fmt"
	"testing"

	qrcode "github.com/skip2/go-qrcode"
)

type QRCodeWriterCalledWith struct {
	Url      string
	Level    qrcode.RecoveryLevel
	Size     int
	Filename string
}
type MockedQRCodeWriter struct {
	CalledWith QRCodeWriterCalledWith
	Called     bool
}

func (m *MockedQRCodeWriter) WriteFile(url string, level qrcode.RecoveryLevel, size int, filename string) error {
	m.CalledWith = QRCodeWriterCalledWith{
		Url:      url,
		Level:    level,
		Size:     size,
		Filename: filename,
	}
	m.Called = true

	return nil
}

func TestNewQRCode(t *testing.T) {
	in := []map[string]string{
		{
			"url":      "url1",
			"filename": "",
		},
		{
			"url":      "url2",
			"filename": "custom1",
		},
		{
			"url":      "url3",
			"filename": "custom2.png",
		},
	}

	want := []map[string]string{
		{
			"url":      "url1",
			"filename": "qrcode.png",
		},
		{
			"url":      "url2",
			"filename": "custom1.png",
		},
		{
			"url":      "url3",
			"filename": "custom2.png",
		},
	}

	for i, v := range in {
		writer := MockedQRCodeWriter{}

		result := NewQRCode(&writer, v["url"], v["filename"])

		if result != want[i]["filename"] {
			t.Error(fmt.Sprintf("Returned wrong filename: %s instead of %s", result, want[i]["filename"]))
		}

		if !writer.Called {
			t.Error("Writer not called")
		}

		if writer.CalledWith.Url != want[i]["url"] {
			t.Error(fmt.Sprintf("Writer called with wrong url: %s instead of %s", writer.CalledWith.Url, want[i]["url"]))
		}

		if writer.CalledWith.Filename != want[i]["filename"] {
			t.Error(fmt.Sprintf("Writer called with wrong filename: %s instead of %s", writer.CalledWith.Filename, want[i]["filename"]))
		}
	}
}
