package wasm

import (
	"slices"
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
	newValue := textValue

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
	} else {
		newValue = rtoolsWasmExports.Get(selValue.String()).Invoke(textValue)
		textArea.Set("value", newValue)
	}

	tEncVal := ""

	if textValue.Type() != js.TypeNull && textValue.Type() != js.TypeUndefined {
		tEncVal = js.Global().Get("encodeURIComponent").Invoke(newValue).String()
	}

	url := js.Global().Get("URL").New(js.Global().Get("location").Get("href"))
	url.Get("searchParams").Call("set", "utility", selValue)
	url.Get("searchParams").Call("set", "text", tEncVal)

	js.Global().Get("history").Call("pushState", nil, "", url.Call("toString"))
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

	if slices.Contains(utilities, utility.String()) {
		el := js.Global().Get("document").Call(
			"querySelector",
			"form select[name=\"utility\"]",
		)

		if el.Type() != js.TypeUndefined {
			el.Set("value", utility)
		}
	}

	if text.Type() != js.TypeNull && text.Type() != js.TypeUndefined {
		el := js.Global().Get("document").Call(
			"querySelector",
			"form textarea[name=\"text\"]",
		)
		tDecVal := js.Global().Get("decodeURIComponent").Invoke(text).String()

		if el.Type() != js.TypeUndefined && el.Type() != js.TypeNull {
			el.Set("value", tDecVal)
		}
	}
	return nil
}
