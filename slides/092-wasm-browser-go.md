
![bg](#3AF1F2)
![fg](#000000)
# Wasm avec Go dans le navigateur

---

<style scoped>
  mark {
    background-color: #EFD217;
    color: #000000;
  }
</style>

# Go + JavaScript = 💖

```bash
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

```html
<script src="wasm_exec.js"></script>
```

 > Disclaimer, I 💛 <mark>**JavaScript**</mark>
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
mark-cyan {
  background-color: #44F0EF;
  color: #000000;
}
ul {
  font-size: 60%;
}
</style>

#### Fonction (<mark-purple>wasm</mark-purple>) en <mark-cyan>Go</mark-cyan>
##### Appelée en <mark>JavaScript</mark>

```go
func Hello(this js.Value, args []js.Value) interface{} {
  message := args[0].String() // get the parameters
  return "😃 Hello " + message
}
```

- 2 paramètres et une `interface en retour`
- le 1er `this` fait référence à l'objet global `window`
- le second est un slice de `[]js.Value` (ensemble des arguments passés lors de l'appel à partir de <mark>JavaScript</mark>)


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
ul {
  font-size: 60%;
}
</style>

#### Initialiser la fonction

```go
func main() {

  js.Global().Set("Hello", js.FuncOf(Hello))

  // make sure that the go program won't exit
  <-make(chan bool)
}
```
- La fonction `Hello` est rattaché à l'objet `Global` de <mark>JavaScript</mark>
- Utilisation d'une `channel` pour éviter "de sortir"

<!--
Et avec ça, on peut faire plein de choses ...
Comme en JavaScript 😉
-->

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
mark-cyan {
  background-color: #44F0EF;
  color: #000000;
}
ul {
  font-size: 60%;
}
</style>

#### Utilisation de la fonction <mark-cyan>Go</mark-cyan> en <mark>JavaScript</mark>

```javascript
const go = new Go() // Go Wasm runtime

WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject)
  .then(resultObject => {
    // execute `main`
    go.run(resultObject.instance)
    // instance object contains
    // all the Exported WebAssembly functions
    let resultValue = Hello("Bob Morane")
    //😃 Hello "Bob Morane
  })
  .catch(error => {
    console.log("😡 ouch", error)
  })
```

<!--
Il est temps de voir quelques exemples
-->

---
### Mais aussi ...
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
mark-cyan {
  background-color: #44F0EF;
  color: #000000;
}
ul {
  font-size: 60%;
}
</style>

###### Appeler une fonction <mark>JavaScript</mark> à partir d'une Fonction (<mark-purple>wasm</mark-purple>) en <mark-cyan>Go</mark-cyan>

```go
import (
	"syscall/js"
)

func main() {

	message := "👋 Hello World from Go 🌍"

	// ! We got a reference to the DOM
	document := js.Global().Get("document")
	h2 := document.Call("createElement", "h2")
	h2.Set("innerHTML", message)
	document.Get("body").Call("appendChild", h2)

}
```

- `"syscall/js"` permet à WebAssembly d'accéder à l'hôte (navigateur)
- la méthode `Call` permet d'appeler des fonctions <mark>JavaScript</mark> (std+udf)

---
<style scoped>
mark-cyan {
  background-color: #44F0EF;
  color: #FFFFFF;
}
</style>

![bg](#000000)
![fg](#FFFFFF)
# Démos 🚀

<a href="https://github.com/wasm-university/sunny-tech-2022/tree/main/01-go-hello" target="_blank">01-go-hello</a>
<a href="https://github.com/wasm-university/sunny-tech-2022/tree/main/02-wasm-go-boids" target="_blank">02-wasm-go-boids (<mark-cyan>with TinyGo</mark-cyan>)</a>
---
