#!/bin/bash

########################
# include the magic
########################
. ../demo-magic.sh

# hide the evidence
clear

# Put your stuff here
pei "dotnet new web -o hello"
pei "cd hello"
pei "dotnet add package Wasi.Sdk --prerelease"
pei "dotnet add package Wasi.AspNetCore.Server.Native --prerelease"

echo "🖐️ update source code"
