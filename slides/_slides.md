---
theme: uncover
size: 16:9
paginate: true

---
<style scoped>
  mark {
    background-color: #942EC1;
    color: #FFFFFF;
  }
</style>
# Petit cours de <mark>Wasm</mark> front/back par l'exemple

**SunnyTech 2022**

https://github.com/wasm-university/sunny-tech-2022

---
# Merci 😃

**à vous & à l'équipe Voxxed Luxembourg**

---

![bg left:40% 80%](pictures/k33g.png)

#### Philippe Charrière

- TAM @GitLab
- 🤖 Bots.Garden
- 🦊 @k33g
- 🐦 @k33g_org
- 🍊🦸Gitpod Hero
- GDG Cloud IOT Lyon
- RdV des speakers

---
# Déroulement

- 👋 Vous pouvez intervenir à tout moment
- 10% Théorie 90% Démos (en gros)
- 🚀 Des démos que vous pourrez refaire :
  - https://github.com/wasm-university
  - en utilisant <mark>Gitpod</mark>
  - ou en utilisant <mark>DevContainer</mark>

---
# Objectif(s)

- Université "découverte" par l’exemple
- Rien de complexe
- Repartir avec le bagage nécessaire 🧳

🖐️ Ne posez pas de questions compliquées 😛🙏
https://github.com/wasm-university/sunny-tech-2022/issues

---
# Une petite remarque ...

- Faire choisir le langage au démarrage 🤔
- En fait **Wasm** est polyglotte
- ... Et j'ai galéré en **Rust** sur les exemples compliqués 🥵

---
![bg](#F0EA71)
# WebAssembly ???

## WASM ???
### C'est parti ! 🚀

---
###### Wasm Quoi/Pourquoi ?

![w:900](pictures/wasm-01.jpeg)


---

###### Histoire

![w:900](pictures/wasm-02.jpeg)

---

## Wasm peut s’exécuter partout

JavaScript (navigateur)
JavaScript (Node.js)
GraalVM
Runtimes **WASI** (Wasmer, Wastime, Wasmedge, …): CLI & Libs
<!-- webassembly system interface -->
---

Wasm file ~= container image, **smaller**, safer, without an OS

---
###### Hosts

![w:900](pictures/wasm-03.jpeg)


<!-- la portabilité de wasm dépend de l'hôte -->

---
![bg](#C4D8F8)
# Wasm a "quelques" limitations

---

### 🖐️ Le module Wasm n’accède pas à l’OS

- Wasm c’est pour du compute (au départ)
- Pas d’accès aux fonctions systèmes de l’OS (hors host functions)
  - I/O
  - Sockets
- Pas d’accès à la mémoire hors allocation spécifique
<!-- vérifier cette partie -->

---

## <mark>Safe by default</mark>

### C'est une bonne limitation
---
### 📣 La Communication Wasm <=> Host  n’est pas triviale
> (trop bas niveau ?)

#### 4 types de données pour les paramètres:

  - 32 & 64 Bit Integer
  - 32 & 64 Bit Floating-Point Number

---
## String 😡

> Mais ...
> - Certains "hôtes" (et toolchains) ont déjà tout prévu
> - Certains frameworks aussi pour WASI *(WebAssembly System Interface)*
---

![bg](#B8F6C5)
# 🛠 ToolChains

---
##### Toolchains par langage & hôte

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
</style>

Langage         | WASM (VM JS)                    | WASI                                     | Remarks
:---------------|:--------------------------------|:-----------------------------------------|:--------
C/C++           | EMScripten, LLVM (clang)        | LLVM, SDK C/C++ Wasmer                   |
<mark>Rust</mark>            | Wasm-pack + wasm-bindgen (glue) | rustup target add wasm32-wasi            | <mark>support navigateur</mark> 💖
<mark>Go</mark>              | Intégré à la toolchain standard | Non ou alors utiliser TinyGo             | <mark>support navigateur</mark> 💖
Assemblyscript  | Intégré                         | Intégré                                  | Ne cible que du WASM
Swift           | SwiftWasm                       | SwiftWasm                                |
Kotlin          | Kotlin native (expérimental)    |                                          |
C#              | Blazor (solution complète)      | <mark>dotnet add package Wasi.Sdk --prerelease</mark> |
Ruby            | Artichoke                       | En cours (portage CRuby Wasm32-WASI)     |
Python          | Expérimental                    |                                          |

<!-- regarder prez de Sébastien pour Kotlin -->
###### *Liste non exhaustive*
---
![bg](#E3C3E9)
# Statut actuel de Wasm

https://blog.scottlogic.com/2022/06/20/state-of-wasm-2022.html

---
![w:900](pictures/wasm-state.png)

---
<style scoped>
mark {
  background-color: #EFD217;
  color: #000000;
}
mark-purple {
  background-color: #942EC1;
  color: #FFFFFF;
}

mark-orange {
  background-color: #F0B044;
  color: #000000;
}

mark-cyan {
  background-color: #44F0EF;
  color: #000000;
}

mark-grey {
  background-color: #E2E0D6;
  color: #000000;
}

mark-green {
  background-color: #71F09C;
  color: #000000;
}

</style>

#### (Très) Rapide résumé (issue du sondage)

- L'utilisation de <mark-purple>**WASM**</mark-purple> ++ fréquente
- Popularité de <mark-orange>**Rust**</mark-orange> en hausse
- De + en + de personnes veulent faire du <mark-purple>**WASM**</mark-purple> en <mark-cyan>**Go**</mark-cyan>
- <mark-green>**Wasmtime**</mark-green> est le runtime le plus utilisé (arm ? 🤔)
- Utilisation de <mark-purple>**WASM**</mark-purple> pour <mark-grey>**Serverless** & **plug-ins**</mark-grey> en hausse
- <mark>**JavaScript**</mark> est devenu un langage utilisable pour <mark-purple>**WASM**</mark-purple> 😮🤔

<!--
Rust usage and desireabillity has continued to climb
Python has seen a big climb in usage
JavaScript has become a viable WebAssembly language
It’s been a good year for Blazor, with a big climb in usage and desire
Wasmtime is the most widely used runtime
The use of WebAssembly for Serverless, Containerisation and as a plug-in host has climbed significantly
Survey respondents are using WebAssembly much more freq
-->

---
##### Near’s JS SDK based on QuickJS

![w:800](pictures/js-wasm-01.png)

---
##### JS to WebAssembly toolchain

![w:800](pictures/js-wasm-02.png)

---
##### Bringing JavaScript and TypeScript to Suborbital

![w:800](pictures/js-wasm-03.png)

---
# Liens relatifs à Wasm 💜 JavaScript

- [thread Twitter de @BrendanEich](https://twitter.com/BrendanEich/status/1535304420426141696?ref_src=twsrc%5Etfw%7Ctwcamp%5Etweetembed%7Ctwterm%5E1535304420426141696%7Ctwgr%5E%7Ctwcon%5Es1_&ref_url=https%3A%2F%2Fblog.scottlogic.com%2F2022%2F06%2F20%2Fstate-of-wasm-2022.html)
- Javy (Shopify) : https://github.com/Shopify/javy
- Suborbital Blog Post : https://blog.suborbital.dev/bringing-javascript-and-typescript-to-suborbital
