package wasm

import (
	"fmt"
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

func callWasm(_ js.Value, args []js.Value) any {
	event := args[0]
	event.Call("preventDefault")

	form := event.Get("srcElement")

	textArea := form.Get("elements").Get("text")
	textValue := textArea.Get("value")

	selectEl := form.Get("elements").Get("utility")
	selValue := selectEl.Get("value")

	img := form.Call("querySelector", "img.qrCode")

	rtoolsWasmExports := js.Global().Get("rtoolsWasmExports")
	img.Set("src", "")
	if selValue.String() == "getQrCode" {
		createObjectURL := js.Global().Get("URL").Get("createObjectURL")

		obj := js.Global().Get("Object").New()
		obj.Set("type", "image/png")
		bytes := rtoolsWasmExports.Get(selValue.String()).Invoke(textValue)
		arr := js.Global().Get("Array").New(3)
		arr.SetIndex(0, bytes)
		blob := js.Global().Get("Blob").New(arr, obj)
		url := createObjectURL.Invoke(blob)

		img.Set("src", url)
		return false
	}

	newValue := rtoolsWasmExports.Get(selValue.String()).Invoke(textValue)
	textArea.Set("value", newValue)

	tEncVal := ""

	if textValue.Type() != js.TypeNull && textValue.Type() != js.TypeUndefined {
		tEncVal = js.Global().Get("encodeURIComponent").Invoke(newValue).String()
	}

	newSearch := fmt.Sprintf("?utility=%s&text=%s", selValue, tEncVal)
	js.Global().Get("location").Set("search", newSearch)
	return false
}

func loadFromUrl(_ js.Value, _ []js.Value) any {
	utilities := []string{
		"base64Encode",
		"base64Decode",
		"jsonBeautify",
		"jsonMinify",
		"generateUuid",
		"getQrCode",
	}
	search := js.Global().Get("location").Get("search")
	params := js.Global().Get("URLSearchParams").New(search)

	utility := params.Call("get", "utility")
	text := params.Call("get", "text")

	fmt.Printf("stuf: %s,  %s, %v\n", utility, text, utilities)
	return nil
}
