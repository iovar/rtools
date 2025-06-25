package main

import (
	"fmt"

	"github.com/iovar/rtools/pkg/wasm"
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

	wasm.Setup()

	select {} // Keeps the program running
}
