package main

import (
	"fmt"
	"syscall/js"

	"github.com/iovar/rtools/pkg/tools"
)

/*
The plan for wasm is to:

Done
1. First establish simple communication between wasm and js
4. Serve a simple html, which simply calls e.g. uuid

Not done
2. Expose the main tools function for uuid first
3. add a simple server started in wasm/main.go
5. add different routes for the different routes, and create some links header
6. replace the header with something injected from the wasm module
7. start on the html wasm, read the route, spit the write html
8. Connect all tools, with forms, etc
*/
func main() {
	fmt.Println("Go WASM initialized")

	fmt.Printf("Go %d\n", callAdd(2, 3))

	js.Global().Set("generateUuid", js.FuncOf(generateUuid))
	js.Global().Set("base64Encode", js.FuncOf(base64Encode))
	js.Global().Set("base64Decode", js.FuncOf(base64Decode))

	select {} // Keeps the program running
}

func callAdd(x, y int) int {
	addFn := js.Global().Get("rtoolsWasmExports").Get("add")
	result := addFn.Invoke(x, y)
	return result.Int()
}

func generateUuid(_ js.Value, _ []js.Value) any {
	return js.ValueOf(tools.NewUuid())
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
