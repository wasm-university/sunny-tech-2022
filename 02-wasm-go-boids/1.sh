#!/bin/bash

########################
# include the magic
########################
. ../demo-magic.sh

# hide the evidence
clear

# Put your stuff here
rm wasm_exec.js
tinygo_version="0.23.0"
#wget https://raw.githubusercontent.com/tinygo-org/tinygo/v${tinygo_version}/targets/wasm_exec.js
pei "tinygo build -o main.wasm -target wasm ./main.go"
pei "ls -lh *.wasm"
echo "🚀 run http server 🌍"
pkill -f little-server
gp url 8081
pe "bash -c \"exec -a little-server-8081 node server.js\""
