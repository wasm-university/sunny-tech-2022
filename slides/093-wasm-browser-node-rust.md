
<style scoped>
  mark {
    background-color: #EFD217;
    color: #000000;
  }
  mark-green {
    background-color: #12984E;
    color: #FFFFFF;
  }
  mark-orange {
    background-color: #F0B044;
    color: #000000;
  }
</style>

![bg](#3AF1F2)
![fg](#000000)

# Wasm avec <mark-orange>Rust</mark-orange> dans le navigateur et aussi <mark-green>Node.js</mark-green>
##### 2 VMs <mark>JavaScript</mark>

## 🦀 + 🕸️ = 💖

https://rustwasm.github.io/

---

# Facile ?
## avec Wasm Bindgen, OUI ‼️ 😍

https://github.com/rustwasm/wasm-bindgen
> Facilitating high-level interactions between Wasm modules and JavaScript

---

#### Créer un projet "Rust Wasm"

###### <mark>Créer un projet de type "library"</mark>

```bash
cargo new --lib hello
```

###### <mark>Mise à jour de `Cargo.toml`</mark>

```toml
[lib]
name = "hello"
path = "src/lib.rs"
crate-type =["cdylib"]

[dependencies]
wasm-bindgen = "0.2.50"
```

---

###### <mark>Modifier `main.rs`<mark>

```rust
use wasm_bindgen::prelude::*;

#[wasm_bindgen]
pub fn hello(s: String) -> String {
  let r = String::from("👋 hello ");

  return r + &s;
}
```

---

###### <mark>Compiler pour le navigateur<mark>

```bash
cd hello
wasm-pack build --release --target web
```
> 🖐️ `--target web`

###### <mark>Compiler pour Node.js<mark>

```bash
wasm-pack build --release --target nodejs
```
> 🖐️ `--target nodejs`

---

######  <mark>Utiliser dans le navigateur<mark>

```html
<script type="module">
  import init, { hello } from './hello/pkg/hello.js'

  async function run() {
    await init()
    console.log(hello("Bob Morane"))
  }
  run();
</script>
```

######  <mark>Utiliser avec Node.js<mark>


```javascript
const wasm = require("./hello/pkg/hello")

console.log(wasm.hello("Bob Morane")
```

---

![bg](#000000)
![fg](#FFFFFF)
# Démos 🚀


<a href="https://github.com/wasm-university/sunny-tech-2022/tree/main/03-nodejs-rust-function" target="_blank">03-nodejs-rust-function</a>

---
