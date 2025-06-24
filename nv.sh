#!/usr/bin/env bash

# Set Go environment variables for WASM
export GOOS=js
export GOARCH=wasm

# Run Neovim with any arguments passed to this script
nvim "$@"
