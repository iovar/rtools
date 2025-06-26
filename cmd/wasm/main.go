package main

import (
	"fmt"

	"github.com/iovar/rtools/pkg/wasm"
)

/*
The plan for wasm is to:

Done
1. First establish simple communication between wasm and js
2. Serve a simple html, which simply calls e.g. uuid
3. add a simple server started in wasm/server.go and add make command for it

Not done
4. move js functions to wasm ones, if possible
5. move html template to wasm & inject with render start command
6. add styles & beautify (as html/template)
*/
func main() {
	fmt.Println("Go WASM initialized")

	wasm.Setup()

	select {} // Keeps the program running
}
