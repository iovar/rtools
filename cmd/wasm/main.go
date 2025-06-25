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
3. add a simple server started in wasm/server.go and add make command for it
4. move js functions to wasm ones, if possible
5. move html template to wasm & inject
6. add styles & beautify (as html/template)
7. replace the header with rtools:[toolname] injected from the wasm module
*/
func main() {
	fmt.Println("Go WASM initialized")

	wasm.Setup()

	select {} // Keeps the program running
}
