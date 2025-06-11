package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/iovar/rtools/pkg/tools"
)

func Start() {
	help := flag.Bool("help", false, "Print help message")
	tool := flag.NewFlagSet("tool", flag.ExitOnError)
	base64 := tool.String("base64", "", "base64 (encode/decode)")
	uuid := tool.Bool("uuid", false, "get a uuid")
	json := tool.String("json", "", "json (beautify/minify)")
	var result string

	flag.Parse()

	if len(os.Args) == 1 || *help {
		fmt.Printf("rtools tool [-base64|-json|-uuid]\n\nUsage of tool:\n  -base64 string\n      base64 (encode/decode)\n  -json string\n      json (beautify/minify)\n  -uuid\n      get a uuid\n")
	} else if os.Args[1] == "tool" {
		tool.Parse(os.Args[2:])

		if *base64 == "encode" {
			result = tools.Base64Encode(os.Args[4])
		} else if *base64 == "decode" {
			result = tools.Base64Decode(os.Args[4])
		} else if *uuid == true {
			result = tools.NewUuid()
		} else if *json == "beautify" {
			result = tools.JSONBeautify(os.Args[4])
		} else if *json == "minify" {
			result = tools.JSONMinify(os.Args[4])
		} else {
			fmt.Printf("Unkown command: %v\n", os.Args[2:])
		}
	} else {
		fmt.Printf("Unkown tool: %s\n", os.Args[1])
		os.Exit(1)
	}

	fmt.Printf("%s\n", result)
}
