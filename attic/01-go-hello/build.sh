#!/bin/bash
GOOS=js GOARCH=wasm go build -o main.wasm

ls -lh *.wasm
