package wasm

import (
	"syscall/js"

	"github.com/iovar/rtools/pkg/tools"
)

func generateUuid(_ js.Value, _ []js.Value) any {
	return tools.NewUuid()
}

func base64Encode(_ js.Value, args []js.Value) any {
	if len(args) != 1 {
		return "Invalid number of arguments"
	}

	input := args[0].String()
	return tools.Base64Encode(input)
}

func base64Decode(_ js.Value, args []js.Value) any {
	if len(args) != 1 {
		return "Invalid number of arguments"
	}

	input := args[0].String()
	return tools.Base64Decode(input)
}

func jsonBeautify(_ js.Value, args []js.Value) any {
	if len(args) != 1 {
		return "Invalid number of arguments"
	}

	input := args[0].String()
	return tools.JSONBeautify(input)
}

func jsonMinify(_ js.Value, args []js.Value) any {
	if len(args) != 1 {
		return "Invalid number of arguments"
	}

	input := args[0].String()
	return tools.JSONMinify(input)
}

func getQrCode(_ js.Value, args []js.Value) any {
	if len(args) != 1 {
		return "Invalid number of arguments"
	}

	input := args[0].String()

	result := tools.NewQRCodeBlob(input)
	bytes := js.Global().Get("Uint8Array").New(len(result))
	js.CopyBytesToJS(bytes, result)

	return bytes
}
