CLI_BINARY_NAME=rtools
WASM_BINARY_NAME=rtools.wasm
CLI_MAIN_PATH=./cmd/cli/main.go
WASM_MAIN_PATH=./cmd/wasm/main.go

# --- Build ----
.PHONY: build
build: build-cli build-wasm

.PHONY: build-cli
build-cli:
	go build -o $(CLI_BINARY_NAME) $(CLI_MAIN_PATH)

.PHONY: build-wasm
build-wasm:
	GOOS=js GOARCH=wasm go build -o $(WASM_BINARY_NAME) $(WASM_MAIN_PATH)

# --- Run ----
.PHONY: run
run: run-cli

.PHONY: run-cli
run-cli:
	go run $(CLI_MAIN_PATH) $(ARGS)

.PHONY: run-wasm
run-wasm: build-wasm
    # TODO - implement a way to run the wasm binary

# --- Clean ---
.PHONY: clean
clean: clean-cli clean-wasm

.PHONY: clean-cli
clean-cli:
	rm -f $(CLI_BINARY_NAME)

.PHONY: clean-wasm
clean-wasm:
	rm -f $(WASM_BINARY_NAME)

# --- Test ---
.PHONY: test
test:
	go test ./...

# --- Format ---
.PHONY: fmt
fmt:
	go fmt ./...
