<style scoped>
  mark {
    background-color: #17EFE7;
    color: #000000;
  }
  ul {
    font-size: 90%;
  }
</style>


#### Mais il y a des limitations 😢

- **Système de type trop simple**
  - <mark>Seulement 4 types numériques</mark> :
    - Integers (32 & 64 bit)
    - Floats (32 & 64 bit)
  - Passer une `String` à une fonction n'est pas trivial 🥵

- **Mode d’exécution fermé du module Wasm**
  - <mark>Pas d’accès à "l’extérieur"</mark> :
    - Pas d’appel http
    - Pas d’accès fichier
    - ...
  - Il est possible de définir des **host functions**
    - mais ce n'est pas simple 🤬

---
<style scoped>
  ul {
    font-size: 90%;
  }
</style>


#### Comment contourner ces limitations
##### "the hard way"

- Passage de **Strings** à une fonction avec **WasmEdge**
  - Pass complex parameters to Wasm functions: https://wasmedge.org/book/en/embed/go/memory.html
- Création de **Host Functions** avec **WasmEdge**
  - https://wasmedge.org/book/en/extend/plugin/hostfunction.html#host-functions

---

###### Principe des host functions

![w:900](pictures/wasm-06.jpeg)

---
<style scoped>
  mark {
    background-color: #17EFE7;
    color: #000000;
  }
</style>
## Mais : à venir

<mark><b>Interface Types</b></mark> : décrire des types de plus haut niveau, éviter les frictions

https://hacks.mozilla.org/2019/08/webassembly-interface-types/

---
