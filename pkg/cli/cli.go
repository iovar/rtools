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
	var result string

	flag.Parse()

	if *base64 == "encode" && len(os.Args) > 3 {
		result = tools.Base64Encode(os.Args[3])
	} else if *base64 == "decode" && len(os.Args) > 3 {
		result = tools.Base64Decode(os.Args[3])
	} else if *uuid == true {
		result = tools.NewUuid()
	} else if *json == "beautify" && len(os.Args) > 3 {
		result = tools.JSONBeautify(os.Args[3])
	} else if *json == "minify" && len(os.Args) > 3 {
		result = tools.JSONMinify(os.Args[3])
	} else {
		flag.PrintDefaults()
	}

	fmt.Printf("%s\n", result)
}
