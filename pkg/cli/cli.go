package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/iovar/rtools/pkg/tools"
)

func Start() {
	base64 := flag.String("base64", "", "base64 (encode/decode)")
	uuid := flag.Bool("uuid", false, "get a uuid")
	json := flag.String("json", "", "json (beautify/minify)")
	qrcode := flag.String("qrcode", "", "qrcode URL FILENAME(opt)")
	var result string

	flag.Parse()

	switch {
	case *base64 == "encode" && len(os.Args) > 3:
		result = tools.Base64Encode(os.Args[3])
	case *base64 == "decode" && len(os.Args) > 3:
		result = tools.Base64Decode(os.Args[3])
	case *uuid == true:
		result = tools.NewUuid()
	case *json == "beautify" && len(os.Args) > 3:
		result = tools.JSONBeautify(os.Args[3])
	case *json == "minify" && len(os.Args) > 3:
		result = tools.JSONMinify(os.Args[3])
	case qrcode != nil && len(os.Args) > 2:
		qrFile := ""
		if len(os.Args) > 3 {
			qrFile = os.Args[3]
		}

		qrWriter := tools.QRCodeWriter{}
		result = tools.NewQRCode(&qrWriter, *qrcode, qrFile)
	default:
		flag.PrintDefaults()
	}

	fmt.Printf("%s\n", result)
}
