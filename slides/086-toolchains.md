##### Toolchains par langage & hôte ⚡️⚡️

<style scoped>
table {
    height: 80%;
    width: 100%;
    font-size: 20px;
    color: green;
}
th {
    color: blue;
}
mark {
  background-color: #EFD217;
  color: #000000;
}
mark-bis {
  background-color: #000000;
  color: #EFD217;
}
</style>

Langage         | WASM (VM JS)                    | WASI                                     | Remarks
:---------------|:--------------------------------|:-----------------------------------------|:--------
<mark>C/C++ </mark>          | EMScripten, LLVM (clang)        | LLVM, SDK C/C++ Wasmer                   |
<mark>Rust</mark>            | Wasm-pack + wasm-bindgen (glue) | rustup target add wasm32-wasi            | <mark>support navigateur</mark> 💖
<mark-bis>Go</mark-bis>              | Intégré à la toolchain standard | Non ou alors utiliser TinyGo             | <mark>support navigateur</mark> 💖
<mark>Assemblyscript</mark>  | Intégré                         | Intégré                                  | Ne cible que du WASM
Swift           | SwiftWasm                       | SwiftWasm                                |
Kotlin          | Kotlin native (expérimental)    |                                          |
C#              | Blazor (solution complète)      | <mark>dotnet add package Wasi.Sdk --prerelease</mark> |
Ruby            | Artichoke                       | En cours (portage CRuby Wasm32-WASI)     |
Python          | Expérimental                    |                                          |

<!-- regarder prez de Sébastien pour Kotlin -->
###### *Liste non exhaustive*

---
