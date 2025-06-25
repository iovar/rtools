package tools

import (
	"fmt"
	"strings"

	qrcode "github.com/skip2/go-qrcode"
)

const QR_SIZE = 256
const FILENAME = "qrcode.png"

type QRCodeWriter struct{}

func (q *QRCodeWriter) WriteFile(url string, level qrcode.RecoveryLevel, size int, filename string) error {
	return qrcode.WriteFile(url, level, size, filename)
}

type QRCodeWriterInterface interface {
	WriteFile(url string, level qrcode.RecoveryLevel, size int, filename string) error
}

func NewQRCode(q QRCodeWriterInterface, url string, customFilename string) string {
	defer func() {
		if er := recover(); er != nil {
			fmt.Printf("Recovered after following error: %v\n", er)
		}
	}()

	filename := FILENAME

	if customFilename != "" {
		filename = customFilename
	}

	if !strings.HasSuffix(filename, ".png") {
		filename = filename + ".png"
	}

	err := q.WriteFile(url, qrcode.Medium, QR_SIZE, filename)

	if err != nil {
		panic(err)
	}

	return filename
}

func NewQRCodeBlob(url string) []byte {
	defer func() {
		if er := recover(); er != nil {
			fmt.Printf("Recovered after following error: %v\n", er)
		}
	}()

	img, err := qrcode.Encode(url, qrcode.Medium, QR_SIZE)

	fmt.Printf("URL %v %v %v\n", url, err, len(img))
	if err != nil {
		panic(err)
	}

	return img
}
