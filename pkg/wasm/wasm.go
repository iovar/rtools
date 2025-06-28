package wasm

import (
	"syscall/js"
)

// wrap functions
// start sets all the functions on the go rtoolsWasmExports object
func Setup() {
	wasmExports := js.Global().Get("rtoolsWasmExports")
	wasmExports.Set("generateUuid", js.FuncOf(generateUuid))
	wasmExports.Set("base64Encode", js.FuncOf(base64Encode))
	wasmExports.Set("base64Decode", js.FuncOf(base64Decode))
	wasmExports.Set("jsonBeautify", js.FuncOf(jsonBeautify))
	wasmExports.Set("jsonMinify", js.FuncOf(jsonMinify))
	wasmExports.Set("getQrCode", js.FuncOf(getQrCode))

	wasmExports.Set("callWasm", js.FuncOf(callWasm))
}
