

## Utiliser le WASI SDK pour .NET Core

- Le **Wasi.Sdk** est expérimental
- Il permet de compiler des projets .NET Core en Wasm
- Il est aussi possible de faire des apps ASP.NET Core

---
<style scoped>
  mark {
    background-color: #17EFE7;
    color: #000000;
  }
</style>

### Application console

```bash
dotnet new console -o hello
cd hello
dotnet add package Wasi.Sdk --prerelease
dotnet build
```

<mark>16M bin/Debug/net7.0/hello.wasm</mark>

```bash
wasmtime bin/Debug/net7.0/hello.wasm
```
---
<style scoped>
  mark {
    background-color: #17EFE7;
    color: #000000;
  }
</style>

### Application ASP.Net

```bash
dotnet new web -o hello
cd hello
dotnet add package Wasi.Sdk --prerelease
dotnet add package Wasi.AspNetCore.Server.Native --prerelease
dotnet build
```

<mark>31M bin/Debug/net7.0/hello.wasm</mark>

```bash
wasmtime bin/Debug/net7.0/hello.wasm --tcplisten localhost:8080
```

---
![bg](#000000)
![fg](#FFFFFF)
# Démos 🚀
## Utilisation du SDK .Net

<a href="https://github.com/wasm-university/training/tree/main/17-dotnet-wasi-cli-app" target="_blank">17-dotnet-wasi-cli-app</a>
<a href="https://github.com/wasm-university/training/tree/main/18-dotnet-wasi-asp" target="_blank">18-dotnet-wasi-asp</a>

<!-- montrer le code -->
---

<style scoped>
  mark {
    background-color: #FFFFFF;
    color: #000000;
  }
</style>
![bg](#FFC300 )
# Perspectives (Wasi & SDK)

- Write once, run anywhere (encore un peu de travail)
  - Runtimes multi-plateformes
- Applications (CLI ou autre) avec plugins wasm
- "Lanceurs/Serveurs" de modules wasm
  - <mark>**Sécurité** 🖐️</mark>
  - Activation/Ajout de fonctionnalités
  - Bots, Hooks, FaaS, UDF, ...


---
<style scoped>
  mark {
    color: #44F099;
  }
</style>

![bg](#1A8B6E)
![fg](#FFFFFF)

# MicroServices, FaaS, ...
### <mark>Avec WebAssembly</mark>

<!--
- Utiliser d'autres langages (ex Grain)
- Parler des tests de charges
-->
---

# Wagi

![w:100](pictures/deislab.png)
https://deislabs.io/

- **WAGI**: WebAssembly Gateway Interface https://github.com/deislabs/wagi
- **Hippo**, the WebAssembly PaaS https://github.com/deislabs/hippo

---

![w:900](pictures/wasm-08.jpeg)


---

```go
package main

import "fmt"

func main() {
	fmt.Println("content-type: text/plain;utf-8")
	fmt.Println("")
	fmt.Println("👋 Hello World 🌍")
}
```
> GET
```bash
curl http://localhost:3000/hello
```
---

# Spin

![w:100](pictures/fermyon.png)
https://www.fermyon.com/

![w:200](pictures/spin.png)
https://spin.fermyon.dev/


---

### Spin
##### L'évolution de Wagi | Rust & Go (mais pas que)

- Micro Service https://www.wasm.builders/fermyon/a-one-line-microservice-with-spin-33ke
- Bartholomew MicroCMS https://github.com/fermyon/bartholomew
- Kitchensink https://github.com/fermyon/spin-kitchensink

---

##### Spin service en Go (TinyGo)

```go
package main

import (
 "fmt"
 "net/http"

 spinhttp "github.com/fermyon/spin/sdk/go/http"
)

func init() {
 spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/plain")
  fmt.Fprintln(w, "Hello Fermyon!")
 })
}

func main() {}
```

---

# Sat 😍

![w:400](pictures/suborbital.png)

---

# Sat
Pas uniquement une librairie, mais aussi un serveur http
- À base de "**Runnables**"
- Multi langages
  - Rust, Go, AssemblyScript, Grain, Swift, JavaScript, TypesSript)
- Facile à utiliser

---

![w:900](pictures/wasm-09.jpeg)


---

##### Sat [Swift]
`subo create runnable yo --lang swift`

```swift
import Suborbital

class Yo: Suborbital.Runnable {
    func run(input: String) -> String {
        return "hello " + input
    }
}

Suborbital.Set(runnable: Yo())
```

---

### Build, Run
#### Serverless function

`subo build yo`
`SAT_HTTP_PORT=8080 sat ./yo/yo.wasm`
`curl -d "Bob Morane" http://localhost:8080`

https://www.wasm.builders/suborbital/a-simple-data-hashing-serverless-function-using-sat-3fn0

---

#### Wapm.io (Wasmer)
Wasm registry

`SAT_HTTP_PORT=8080 sat https://registry-cdn.wapm.io/contents/k33g/forty-two/1.0.0/forty-two.wasm`

https://www.wasm.builders/k33g_org/publish-your-runnables-on-wapmio-49k0

---

#### Fly.io
Caas | Déploiement ultra simple

https://www.wasm.builders/k33g_org/deploy-a-sat-serverless-function-with-to-flyio-35df

---

# Wasm Cloud

![](pictures/wasmcloud.png)
![](pictures/cosmonic.png)

https://cosmonic.com/

---

# Wasm Cloud
https://cosmonic.com/ | https://wasmcloud.com/

Une plateforme + un SDK + un Runtime
- Nats (communications) https://docs.nats.io/
- Système à base d'acteurs

---
<style scoped>
code {
   font-size: 60%;
}
</style>

##### Wasm Cloud [Rust]

`wash new actor hello --template-name hello`

```rust
/// Implementation of HttpServer trait methods
#[async_trait]
impl HttpServer for HelloActor {

    async fn handle_request(&self, _ctx: &Context, req: &HttpRequest,) -> std::result::Result<HttpResponse, RpcError> {
        let text = ...;

        Ok(HttpResponse {
            body: format!("Hello {}", text).as_bytes().to_vec(),
            ..Default::default()
        })
    }
}
```

---

<style scoped>
code {
   font-size: 60%;
}
</style>

##### Wasm Cloud [TinyGo]

`wash new actor hello --template-name echo-tinygo`

```go
func main() {
	me := Hey{}
	actor.RegisterHandlers(httpserver.HttpServerHandler(&me))
}

type Hey struct{}

func (e *Hey) HandleRequest(ctx *actor.Context, req httpserver.HttpRequest) (*httpserver.HttpResponse, error) {
	r := httpserver.HttpResponse{
		StatusCode: 200,
		Header:     make(httpserver.HeaderMap, 0),
		Body:       []byte("hello"),
	}
	return &r, nil
}
```
---

<style scoped>
  mark {
    background-color: #F7C00E;
    color: #000000;
  }
</style>

![bg](#000000)
![fg](#FFFFFF)
# Démo(s) 🚀

<a href="https://github.com/wasm-university/training/tree/main/22-wagi" target="_blank">22-wagi</a>
<a href="https://github.com/wasm-university/training/tree/main/23-spin" target="_blank">23-spin</a>
<a href="https://github.com/wasm-university/training/tree/main/24-sat" target="_blank">24-sat</a>

---

<style scoped>
  mark {
    color: #44F099;
  }
</style>

![bg](#1A8B6E)
![fg](#FFFFFF)

# Et après ?
### <mark>Le futur de WebAssembly</mark> 👀


---
<style scoped>
ul {
   font-size: 70%;
}
</style>
##### À suivre https://github.com/WebAssembly
*Juin 2022 :*
- https://github.com/WebAssembly/component-model
  - amélioration intégration host
  - activité (GitHub Insights) en début d'année et un peu récemment
- https://github.com/WebAssembly/interface-types
  - reporté dans component-model
- https://github.com/WebAssembly/exception-handling
  - grosse activité en début d'année
- https://github.com/WebAssembly/debugging
  - 💀 😢
- https://github.com/WebAssembly/wasi-filesystem
  - un peu d'activité récemment


---
<style scoped>
  mark {
    color: #44F099;
  }
</style>

![bg](#1A8B6E)
![fg](#FFFFFF)

# Merci 😃

## <mark>Questions ?<mark>
---
