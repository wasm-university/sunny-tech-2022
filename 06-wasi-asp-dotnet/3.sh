#!/bin/bash

########################
# include the magic
########################
. ../demo-magic.sh

# hide the evidence
clear

# Put your stuff here
pei "cd hello"
pei "wasmtime bin/Debug/net7.0/hello.wasm --tcplisten localhost:8080"


