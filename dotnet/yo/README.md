dotnet new web -o yo

cd yo
dotnet add package Wasi.Sdk --prerelease
dotnet add package Wasi.AspNetCore.Server.Native --prerelease
dotnet build


wasmtime bin/Debug/net7.0/yo.wasm --tcplisten localhost:8080

ASP.NET Core Application on Sat

https://youtu.be/A0vz_BWxIMc?t=1506