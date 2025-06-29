CLI_BINARY_NAME=rtools
WASM_BINARY_NAME=rtools.wasm
CLI_MAIN_PATH=./cmd/cli/main.go
SERVER_MAIN_PATH=./cmd/server/main.go
WASM_MAIN_PATH=./cmd/wasm/main.go
WEB_PATH=./res/web
DIST_PATH=./dist

# --- Build ----
.PHONY: build
build: build-cli build-wasm

.PHONY: build-cli
build-cli:
	go build -o $(CLI_BINARY_NAME) $(CLI_MAIN_PATH)

.PHONY: build-wasm
build-wasm: setup
	GOOS=js GOARCH=wasm go build -o $(DIST_PATH)/$(WASM_BINARY_NAME) $(WASM_MAIN_PATH)

# --- Setup ---
.PHONY: setup
setup:
	mkdir -p $(DIST_PATH)
	cp $(WEB_PATH)/*.html $(DIST_PATH)/
	cp $(WEB_PATH)/*.css $(DIST_PATH)/
	cp $$(go env GOROOT)/lib/wasm/wasm_exec.js $(DIST_PATH)/wasm_exec.js
	chmod +w $(DIST_PATH)/wasm_exec.js

.PHONY: compress-wasm
compress-wasm:
	gzip -9 --keep -f $(DIST_PATH)/$(WASM_BINARY_NAME)
	gzip -9 --keep -f $(DIST_PATH)/wasm_exec.js

# --- Dist ---
.PHONY: dist
dist: build-wasm compress-wasm

.PHONY: dist-release
dist-release: dist
	rm -rf ./docs/*
	cp -r $(DIST_PATH)/* ./docs/

# --- Run ----
.PHONY: run
run: run-cli

.PHONY: run-cli
run-cli:
	go run $(CLI_MAIN_PATH) $(ARGS)

.PHONY: run-server
run-server: dist
	go run $(SERVER_MAIN_PATH) 

# --- Clean ---
.PHONY: clean
clean: clean-cli clean-wasm

.PHONY: clean-cli
clean-cli:
	rm -f $(CLI_BINARY_NAME)

.PHONY: clean-wasm
clean-wasm:
	rm -rf $(DIST_PATH)

# --- Test ---
.PHONY: test
test:
	go test ./...

# --- Format ---
.PHONY: fmt
fmt:
	go fmt ./...
