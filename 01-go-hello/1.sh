#!/bin/bash

########################
# include the magic
########################
. ../demo-magic.sh

# hide the evidence
clear

# Put your stuff here
pei "GOOS=js GOARCH=wasm go build -o main.wasm"
pei "ls -lh *.wasm"

# brew install pv
