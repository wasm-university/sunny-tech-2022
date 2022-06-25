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

> Certains "hôtes" (et toolchains) ont déjà tout prévu (certains frameworks aussi pour WASI)
---
